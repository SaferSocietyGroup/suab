package config

import (
	"net/url"
	"errors"
	"strconv"
)

type MasterUrl struct {
	protocol string
	host string
	port uint16
}
func (a *MasterUrl) ToString() string {
	return a.protocol + "://" + a.host + ":" + strconv.Itoa(int(a.port))
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

