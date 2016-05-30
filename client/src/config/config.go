package config

import (
	"os"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"shutupflags"
	"strings"
)

type Config struct {
	DockerImageTag string         `json:"dockerImageTag"`
	MasterUrl string              `json:"masterUrl"`
	SwarmUri string               `json:"swarmUri"`
	Environment map[string]string `json:"environment"`
}

func ReadConfigFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Unable to open config file \"" + path + "\". " + err.Error())
	}

	decoder := json.NewDecoder(file)
	conf := Config{}
	err = decoder.Decode(&conf)
	if err == nil {
		return &conf, nil
	} else {
		return nil, errors.New("Unable to parse config file \"" + path + "\". " + err.Error())
	}
}

func ParseConfigFlags() (*Config, error) {
	dockerImageTag := shutupflags.AddFlag("-d", "--dockerImageTag", "", "The tag of the Docker image in which to build")
	masterUrl := shutupflags.AddFlag("-m", "--master", "", "The SUAB masters URL in the form http://example.com:8080")
	swarmUri := shutupflags.AddFlag("-s", "--swarm", "", "The Docker swarm URI in the form example.com:4000")
	envRaw := shutupflags.AddFlag("-e", "--env", "", "Any environment variables you want in the docker image. Passed as --env a=b,c=d")

	flag.Usage = func () {
		fmt.Println(shutupflags.Usage())
	}

	flag.Parse()

	env, envErr := parseEnv(*envRaw)

	errs := make([]string, 0)
	if envErr != nil {
		errs = append(errs, envErr.Error())
	}
	errorRows := strings.Join(errs, "\n")

	if len(errorRows) > 0 {
		return nil, errors.New(errorRows)
	} else {
		return &Config{
			DockerImageTag: *dockerImageTag,
			MasterUrl: *masterUrl,
			SwarmUri: *swarmUri,
			Environment: env,
		}, nil
	}
}

func parseEnv(raw string) (map[string]string, error) {
	if len(raw) > 0 {
		env, err := stringToEnv(raw)

		if err != nil {
			return nil, err
		} else if len(env) == 0 {
			// there was some attempt to set env vars, but we got none...
			return nil, errors.New("Failed to parse the environment flag")
		}
		return env, nil
	}
	return nil, nil
}

func stringToEnv(raw string) (map[string]string, error) {
	parts := strings.Split(raw, ",")

	toReturn := make(map[string]string)
	for _, rawEnv := range parts {
		envParts := strings.Split(rawEnv, "=")
		if len(envParts) != 2 {
			return nil, errors.New("The environment variables seem malformed, see " + rawEnv)
		}
		// The following overwrites vars if they're specified more than once. Is that good?
		toReturn[envParts[0]] = envParts[1]
	}

	return toReturn, nil
}

func ReadAndParseEffectiveConf(configFilePath string) (*Config, error){
	flagsConf, err := ParseConfigFlags()
	if err != nil {
		return nil, err
	}

	if fileExists(configFilePath) {
		fileConf, err := ReadConfigFile(configFilePath)
		if err != nil {
			return nil, err
		}
		return Merge(*flagsConf, *fileConf), nil
	} else {
		return flagsConf, nil
	}
}
func fileExists(path string) bool {
	_, err := os.Stat(path);
	return err == nil
}

func Merge(important Config, lessImportant Config) *Config {
	if len(important.DockerImageTag) == 0 {
		important.DockerImageTag = lessImportant.DockerImageTag
	}

	if important.MasterUrl == "" {
		important.MasterUrl = lessImportant.MasterUrl
	}

	if important.SwarmUri == "" {
		important.SwarmUri = lessImportant.SwarmUri
	}

	if important.Environment == nil {
		// TODO: Should be done on a var-per-var basis? Either --env overwrites all vars, or it appends new and overwrites some.
		important.Environment = lessImportant.Environment
	}

	return &important
}
