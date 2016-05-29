package main

import (
    "github.com/gin-gonic/gin"
    "os"
    "io"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "errors"
    "log"
    "net/http"
    "path"
)

type Build struct {
    Id string
    Image string
    Host string
    EnvVars map[string]string
    Logs string
}

var buildDir string = "builds"

func main() {
    var fileServer http.Handler = http.FileServer(assetFS())

    r := gin.Default()
    r.Use(CORSMiddleware())

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.POST("/build/:buildId", CreateBuild)
    r.GET("/build/:buildId", GetBuildMetadata)
    r.GET("/builds", ListBuilds)
    r.POST("/build/:buildId/logs", WriteLogs)
    r.GET("/build/:buildId/logs", GetLogs)
    r.POST("/build/:buildId/artifacts/:artifactId", WriteArtifact)
    r.GET("/build/:buildId/artifacts/:artifactId", GetArtifact)
    r.GET("/build/:buildId/artifacts", ListArtifacts)

    setupWebUiEndpoints(r, fileServer)
    setupClientDownloads(r, fileServer)

    r.Run() // listen and server on 0.0.0.0:8080
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func setupWebUiEndpoints(r *gin.Engine, fileServer http.Handler) {
    r.GET("/", func(c *gin.Context) {
        c.Redirect(301, "/web-ui")
    })
    r.Any("/web-ui/*any", func(c *gin.Context) {
        fileServer.ServeHTTP(c.Writer, c.Request)
    })

}

func setupClientDownloads(r *gin.Engine, fileServer http.Handler) {
    r.GET("/client/linux", func(c *gin.Context) {
        c.Redirect(301, "/client/build/suab")
    })
    r.GET("/client/win", func(c *gin.Context) {
        c.Redirect(301, "/client/build/suab.exe")
    })
    r.GET("/client/build/*any", func(c *gin.Context) {
        fileServer.ServeHTTP(c.Writer, c.Request)
    })

}

func CreateBuild(c *gin.Context) {
    buildId := c.Param("buildId")
    if len(buildId) == 0 {
        c.String(400, "You must specify a build id")
        return
    }

    err := os.RemoveAll(path.Join(buildDir, buildId))
    if err != nil {
        log.Printf("Unable to clean the build directory for build %s, %s\n", err, buildId)
    }
    err = os.MkdirAll(path.Join(buildDir, buildId), 0777)
    if err != nil {
        log.Printf("Unable to create build directory for build %s, %s\n", buildId, err.Error())
        c.String(500, "Unable to create build directory")
    }

    WriteFile(path.Join(buildDir, buildId, "metadata"), c.Request.Body)
    c.String(200, "Build created successfully")
}

func WriteFile(file string, source io.Reader) error {
    out, err := os.Create(file)
    defer out.Close()
    if err != nil {
       return err
    }

    written, err := io.Copy(out, source)
    if(err == nil) {
        fmt.Printf("Written: %d", written)
        return nil
    } else {
        return errors.New("Unable to write file " + file + ", " + err.Error())
    }
}

func GetBuildMetadata(c *gin.Context) {
    id := c.Param("buildId")
    if len(id) == 0 {
        c.String(400, "You must specify a build id")
        return
    }

    data, err := ioutil.ReadFile(path.Join(buildDir, id, "metadata"))
    if err == nil {
        c.String(200, string(data))
    } else {
        log.Printf("Unable to read meta data for build %s, %s\n", id, err.Error())
        c.String(500, "Unable to read meta data for build " + id + ", " + err.Error())
    }
}

func ListBuilds(c *gin.Context) {
    if _, err := os.Stat(buildDir); os.IsNotExist(err) {
        c.JSON(200, make([]interface{}, 0))
        return
    }

    files, err := ioutil.ReadDir(buildDir)
    if err != nil {
        log.Printf("Unable to list builds %s\n", err)
        c.String(500, "Failed listing builds %s", err)
        return
    }

    a := make(map[string]interface{}, 0)
    for _, f := range files {
        if f.IsDir() {
            data, err := ioutil.ReadFile(path.Join(buildDir, f.Name(), "metadata"))
            if err == nil {
                var js interface{}
                err = json.Unmarshal(data, &js)
                if err == nil {
                    a[f.Name()] = js
                } else {
                    log.Printf("Could not parse the metadata for build %s as JSON, %s\n", f.Name(), err)
                }
            } else {
                log.Printf("Could not read the metadata for build %s, %s\n", f.Name(), err)
            }
        }
    }
    // TODO: Return 500 if any errors occurred in the loop above
    c.JSON(200, a)
}

func WriteLogs(c *gin.Context) {
    id := c.Param("buildId")
    if len(id) == 0 {
        c.String(400, "You must specify a build id")
        return
    }

    err := WriteFile(path.Join(buildDir, id, "logs"), c.Request.Body)
    if err == nil {
        c.String(200, "logs written")
    } else {
        log.Printf("Could not write the log file for build %s, %s\n", id, err)
        c.String(500, "Could not write the log file for build %s, %s", id, err)
    }
}

func GetLogs(c *gin.Context) {
    id := c.Param("buildId")
    if len(id) == 0 {
        c.String(400, "You must specify a build id")
        return
    }

    data, err := ioutil.ReadFile(path.Join(buildDir, id, "logs"))
    if err == nil {
        c.String(200, string(data))
    } else {
        log.Printf("Could not read the log file for build %s, %s\n", id, err)
        c.String(500, "Could not read the log file for build %s, %s", id, err)
    }
}

func WriteArtifact(c *gin.Context) {
    buildId := c.Param("buildId")
    artifactId := c.Param("artifactId")
    if len(buildId) == 0 || len(artifactId) == 0 {
        c.String(400, "You must specify both a build id and an artifact id")
        return
    }

    err := os.MkdirAll(path.Join(buildDir, buildId, "artifacts"), 0777)
    if err != nil {
        log.Printf("Could not create artifacts folder for build %s, %s\n", buildId, err)
        c.String(500, "Could not create artifacts folder for build %s, %s", buildId, err)
        return
    }

    err = WriteFile(path.Join(buildDir, buildId, "artifacts", artifactId), c.Request.Body)
    if err == nil {
        c.String(200, "Artifact written")
    } else {
        log.Printf("Failed writing artifact for build %s, %s\n", buildId, err)
        c.String(500, "Failed writing artifact for build %s, %s", buildId, err)
    }
}

func GetArtifact(c *gin.Context) {
    buildId := c.Param("buildId")
    artifactId := c.Param("artifactId")
    if len(buildId) == 0 || len(artifactId) == 0 {
        c.String(400, "You must specify both a build id and an artifact id")
        return
    }

    c.File(path.Join(buildDir, buildId, "artifacts", artifactId))
}

func ListArtifacts(c *gin.Context) {
    buildId := c.Param("buildId")
    if len(buildId) == 0 {
        c.String(400, "You must specify a build id")
        return
    }


    // TODO: Make sure builds/buildId exists
    artifactsDir := path.Join(buildDir, buildId, "artifacts")
    if _, err := os.Stat(artifactsDir); os.IsNotExist(err) {
        c.JSON(200, make([]string, 0))
        return
    }

    files, err := ioutil.ReadDir(artifactsDir)
    if err != nil {
        log.Printf("Could not read artifact folder for build %s, %s\n", buildId, err)
        c.String(500, "Could not read artifact folder for build %s, %s", buildId, err)
        return
    }

    a := make([]string, 0)
    for _, f := range files {
        a = append(a, f.Name())
    }
    c.JSON(200, a)
}
