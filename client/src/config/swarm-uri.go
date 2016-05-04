package config

import (
	"strings"
	"errors"
	"strconv"
)

type SwarmUri struct {
	host string
	port uint16
}
func (a *SwarmUri) ToString() string {
	return a.host + ":" + strconv.Itoa(int(a.port))
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

