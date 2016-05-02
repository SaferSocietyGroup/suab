package config

import (
	"os"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"net/url"
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
	DockerImageTag string
	MasterUrl MasterUrl
	SwarmUri SwarmUri
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
/*	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "docker-image-tag, d",
			Usage: "Build in the Docker image `TAG` ",
		},
		cli.StringFlag{
			Name: "master, m",
			Usage: "The url of the suab master. e.g. http://example.com:8080 ",
		},
		cli.StringFlag{
			Name: "swarm, s",
			Usage: "The uri of the Docker swarm to submit the work to, e.g. example.com:4000 ",
		},
	}
*/


	return nil, nil
}