package main

import (
	"fmt"
	"submitters"
	"config"
	"os"
	"strings"
)

func main() {
	configFilePath := "./.suab-config"
	conf, err := config.ReadAndParseEffectiveConf(configFilePath);
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	if errs := validate(conf, configFilePath); len(errs) > 0  {
		fmt.Printf("Invalid config.\n  %s\n", strings.Join(errs, "\n  "))
		os.Exit(2)
	}

	submitter := submitters.GetSubmitter()
	err = submitter(conf.DockerImageTag, conf.MasterUrl.ToString(), conf.SwarmUri.ToString())
	if err == nil {
		fmt.Println("Successfully shut up and built!")
	} else {
		fmt.Printf("Submission failed. %s\n", err)
	}
}

func validate(conf *config.Config, configFile string) []string {
	errs := make([]string, 0)

	if len(conf.DockerImageTag) == 0 {
		errs = append(errs, "You must specify a docker image tag, either via the -d flag, or via the \"dockerImageTag\": \"some-tag\" in " + configFile)
	}

	if conf.MasterUrl == nil || !conf.MasterUrl.IsValid() {
		errs = append(errs, "You must specify a valid master url, either via the -m flag, or via the \"masterUrl\": \"http://example.com:8080\" in " + configFile)
	}

	if conf.SwarmUri == nil || !conf.SwarmUri.IsValid() {
		errs = append(errs, "You must specify a valid docker swarm uri, either via the -d flag, or via the \"swarmUri\": \"example.com:4000\" in " + configFile)
	}

	return errs
}


