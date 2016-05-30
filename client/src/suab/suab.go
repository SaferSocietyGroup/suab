package main

import (
	"fmt"
	"submitters"
	"config"
	"os"
	"strings"
	"net"
	"errors"
)

func main() {
	// TODO: suab list, list all builds
	// TODO: suab logs BUILD_ID, show the logs from the build BUILD_ID

	conf := getAndValidateConfigOrExit("./.suab.json")
	suabShellScript, err := getSuabShellScript()
	if err != nil {
		fmt.Printf("Could not extract the script to run in the docker container, %v\n", err)
		os.Exit(3)
	}
	baseUrl := conf.MasterUrl + "/build/$SUAB_BUILD_ID"
	suabShellScript = injectVariables(suabShellScript, baseUrl, conf.DockerImageTag)

	submitter := submitters.GetSubmitter()
	exitCode, err := submitter(suabShellScript, conf.DockerImageTag, conf.MasterUrl, conf.SwarmUri, conf.Environment)
	if err == nil {
		if exitCode == 0 {
			fmt.Println("Successfully shut up and built!")
		} else {
			// Exit with the same exit code as the submitter
			os.Exit(exitCode)
		}
	} else {
		fmt.Printf("Submission failed. %s\n", err)
	}
}

func getAndValidateConfigOrExit(configFilePath string) *config.Config {
	externalIP, err := externalIP()
	var defaults config.Config;
	if err == nil {
		defaults = config.Config {
			MasterUrl: externalIP + ":8081",
		}
	} else {
		defaults = config.Config{}
	}

	conf, err := config.ReadAndParseEffectiveConf(configFilePath);
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	conf = config.Merge(*conf, defaults)

	if errs := validate(conf, configFilePath); len(errs) > 0  {
		fmt.Printf("Invalid config.\n  %s\n", strings.Join(errs, "\n  "))
		fmt.Println("\nUse the -h flag for more help")
		os.Exit(2)
	}

	return conf
}

// shamelessly stolen from http://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
// Thanks Sebastian and IanB
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func validate(conf *config.Config, configFile string) []string {
	errs := make([]string, 0)

	if len(conf.DockerImageTag) == 0 {
		errs = append(errs, "You must specify a docker image tag, either via the -d flag, or via the \"dockerImageTag\": \"some-tag\" in " + configFile)
	}

	if conf.MasterUrl == "" {
		errs = append(errs, "You must specify a valid master url, either via the -m flag, or via the \"masterUrl\": \"http://example.com:8080\" in " + configFile)
	}

	return errs
}

func getSuabShellScript() (string, error) {
	data, err := Asset("src/asssets/docker-cmd.sh")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func injectVariables(script string, baseUrl string, imageTag string) string {
	script = strings.Replace(script, "$1", baseUrl, 1)
	script = strings.Replace(script, "$2", imageTag, 1)

	return script
}

