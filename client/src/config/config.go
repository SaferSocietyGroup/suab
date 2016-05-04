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
	MasterUrl *MasterUrl          `json:"masterUrl"`
	SwarmUri *SwarmUri            `json:"swarmUri"`
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
	masterRaw := shutupflags.AddFlag("-m", "--master", "", "The SUAB masters URL in the form http://example.com:8080")
	swarmRaw := shutupflags.AddFlag("-s", "--swarm", "", "The Docker swarm URI in the form example.com:4000")
	envRaw := shutupflags.AddFlag("-e", "--env", "", "Any environment variables you want in the docker image. Passed as --env a=b,c=d")

	flag.Usage = func () {
		fmt.Println(shutupflags.Usage())
	}

	flag.Parse()

	masterUrl, masterErr := parseMasterUrl(*masterRaw)
	swarmUri, swarmErr := parseSwarmUri(*swarmRaw)
	env, envErr := parseEnv(*envRaw)

	errs := make([]string, 0)
	if masterErr != nil {
		errs = append(errs, masterErr.Error())
	}
	if swarmErr != nil {
		errs = append(errs, swarmErr.Error())
	}
	if envErr != nil {
		errs = append(errs, envErr.Error())
	}
	errorRows := strings.Join(errs, "\n")

	if len(errorRows) > 0 {
		return nil, errors.New(errorRows)
	} else {
		return &Config{
			DockerImageTag: *dockerImageTag,
			MasterUrl: masterUrl,
			SwarmUri: swarmUri,
			Environment: env,
		}, nil
	}
}

func parseMasterUrl(raw string) (*MasterUrl, error) {
	if len(raw) > 0 {
		masterUrl, err := stringToMasterUrl(raw)
		if err != nil {
			return nil, errors.New("Unable to parse the master URL. " + err.Error())
		}

		return masterUrl, nil
	}
	return nil, nil
}

func parseSwarmUri(raw string) (*SwarmUri, error) {

	if len(raw) > 0 {
		var err error
		swarmUri, err := stringToSwarmUri(raw)
		if err != nil {
			return nil, errors.New("Unable to parse the swarm URI. " + err.Error())
		}
		return swarmUri, nil
	}
	return nil, nil
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
		return merge(flagsConf, fileConf), nil
	} else {
		return flagsConf, nil
	}
}
func fileExists(path string) bool {
	_, err := os.Stat(path);
	return err == nil
}

func merge(important *Config, lessImportant *Config) *Config {
	if len(important.DockerImageTag) == 0 {
		important.DockerImageTag = lessImportant.DockerImageTag
	}

	if important.MasterUrl == nil {
		important.MasterUrl = lessImportant.MasterUrl
	}

	if important.SwarmUri == nil {
		important.SwarmUri = lessImportant.SwarmUri
	}

	return important
}