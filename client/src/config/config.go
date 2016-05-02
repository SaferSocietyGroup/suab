package config

import (
	"os"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"net/url"
	"flag"
)

type MasterUrl struct {
	protocol string
	host string
	port uint16
}
func (a *MasterUrl) ToString() string {
	return a.protocol + "://" + a.host + ":" + string(a.port)
}
func (a *MasterUrl) IsValid() bool {
	return len(a.protocol) > 0 && len(a.host) > 0 && a.port > 0
}
func stringToMasterUrl(s string) (*MasterUrl, error) {
	url, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	hostParts, err := stringToSwarmUri(url.Host)
	if err != nil {
		return nil, errors.New("The host-port part of the URL must be on the form example.com:8080")
	}

	return &MasterUrl{
		protocol: url.Scheme,
		host: hostParts.host,
		port: hostParts.port,
	}, nil
}


type SwarmUri struct {
	host string
	port uint16
}
func (a *SwarmUri) ToString() string {
	return a.host + ":" + string(a.port)
}
func (a *SwarmUri) IsValid() bool {
	return len(a.host) > 0 && a.port > 0
}
func stringToSwarmUri(s string) (*SwarmUri, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return nil, errors.New("You must specify the swarm uri on the form HOST:PORT")
	}

	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	if port < 1 || port > 65535 {
		return nil, errors.New("Invalid port")
	}

	return &SwarmUri{
		host: parts[0],
		port: uint16(port),
	}, nil
}

type Config struct {
	DockerImageTag string `json:"dockerImageTag"`
	MasterUrl *MasterUrl  `json:"masterUrl"`
	SwarmUri *SwarmUri    `json:"swarmUri"`
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
	dockerImageTag := flag.String("d", "", "The tag of the Docker image in which to build")
	masterRaw := flag.String("m", "", "The SUAB masters URL in the form http://example.com:8080")
	swarmRaw := flag.String("s", "", "The Docker swarm URI in the form example.com:4000")

	flag.Parse()

	var masterUrl *MasterUrl = nil
	if len(*masterRaw) > 0 {
		var err error
		masterUrl, err = stringToMasterUrl(*masterRaw)
		if err != nil {
			return nil, errors.New("Unable to parse the master URL. " + err.Error())
		}
	}

	var swarmUri *SwarmUri = nil
	if len(*swarmRaw) > 0 {
		var err error
		swarmUri, err = stringToSwarmUri(*swarmRaw)
		if err != nil {
			return nil, errors.New("Unable to parse the swarm URI. " + err.Error())
		}
	}

	return &Config{
		DockerImageTag: *dockerImageTag,
		MasterUrl: masterUrl,
		SwarmUri: swarmUri,
	}, nil
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