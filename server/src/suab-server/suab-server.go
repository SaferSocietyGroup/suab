package main

import (
    "github.com/gin-gonic/gin"
    "os"
    "io"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type Build struct {
    Id string
    Image string
    Host string
    EnvVars map[string]string
    Logs string
}

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.POST("/build/:buildId", func(c *gin.Context) {
        CreateBuild(c)
    })
    r.GET("/build/:buildId", func(c *gin.Context) {
        GetBuildMetadata(c)
    })
    r.GET("/builds", func(c *gin.Context) {
        ListBuilds(c)
    })
    r.POST("/build/:buildId/logs", func(c *gin.Context) {
        WriteLogs(c)
    })
    r.GET("/build/:buildId/logs", func(c *gin.Context) {
        GetLogs(c)
    })
    r.POST("/build/:buildId/artifacts/:artifactId", func(c *gin.Context) {
        WriteArtifact(c)
    })
    r.GET("/build/:buildId/artifacts/:artifactId", func(c *gin.Context) {
        GetArtifact(c)
    })
    r.GET("/build/:buildId/artifacts", func(c *gin.Context) {
        ListArtifacts(c)
    })

    r.Run() // listen and server on 0.0.0.0:8080
}

func CreateBuild(c *gin.Context) {
    id := c.Param("buildId")
    os.RemoveAll("builds/"+id)
    os.MkdirAll("builds/"+id, 0777)

    WriteFile("builds/"+id+"/metadata", c.Request.Body)
    c.String(200, "")
}

func WriteFile(file string, source io.Reader) {
    out, _ := os.Create(file)
    defer out.Close()
    written, err := io.Copy(out, source)
    if(err != nil) {
        fmt.Printf("%v\n", err)
    } else {
        fmt.Printf("Written: %d", written)
    }
}

func GetBuildMetadata(c *gin.Context) {
    id := c.Param("buildId")
    data,_ := ioutil.ReadFile("builds/"+id+"/metadata")
    c.String(200, string(data))
}

func ListBuilds(c *gin.Context) {
    files,_ := ioutil.ReadDir("builds")
    a := make(map[string]interface{}, 0)
    for _, f := range files {
        if f.IsDir() {
            data, _ := ioutil.ReadFile("builds/"+f.Name()+"/metadata")
            var js interface{}
            json.Unmarshal(data, &js)
            a[f.Name()] = js
        }
    }
    c.JSON(200, a)
}

func WriteLogs(c *gin.Context) {
    id := c.Param("buildId")
    WriteFile("builds/"+id+"/logs", c.Request.Body)
    c.String(200, "")
}

func GetLogs(c *gin.Context) {
    id := c.Param("buildId")
    data,_ := ioutil.ReadFile("builds/"+id+"/logs")
    c.String(200, string(data))
}

func WriteArtifact(c *gin.Context) {
    buildId := c.Param("buildId")
    artifactId := c.Param("artifactId")

    os.MkdirAll("builds/"+buildId+"/artifacts/", 0777)
    WriteFile("builds/"+buildId+"/artifacts/"+artifactId, c.Request.Body)
}

func GetArtifact(c *gin.Context) {
    buildId := c.Param("buildId")
    artifactId := c.Param("artifactId")

    c.File("builds/"+buildId+"/artifacts/"+artifactId)
}

func ListArtifacts(c *gin.Context) {
    buildId := c.Param("buildId")
    files,_ := ioutil.ReadDir("builds/"+buildId+"/artifacts")

    a := make([]string, 0)
    for _, f := range files {
        a = append(a, f.Name())
    }
    c.JSON(200, a)
}