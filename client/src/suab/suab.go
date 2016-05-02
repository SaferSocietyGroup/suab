package main

import (
	"fmt"
	"submitters"
	"config"
	"os"
	"strings"
)

func main() {
	// I want to run suab with no arguments to build the project I'm in
	// The following three variables should come from a config file 
	// and be overridable by flags

	a, err := config.ParseConfigFlags();
	fmt.Printf("a: %+v, err: %+v\n", a, err)

	configFilePath := "./.suab-config"
	conf := &config.Config{

	}

	if errs := validate(conf, configFilePath); len(errs) > 0  {
		fmt.Printf("Invalid config.\n  %s\n", strings.Join(errs, "\n  "))
		os.Exit(1)
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

	if !conf.MasterUrl.IsValid() {
		errs = append(errs, "You must specify a valid master url, either via the -m flag, or via the \"masterUrl\": \"http://example.com:8080\" in " + configFile)
	}

	if !conf.SwarmUri.IsValid() {
		errs = append(errs, "You must specify a valid docker swarm uri, either via the -d flag, or via the \"swarmUri\": \"example.com:4000\" in " + configFile)
	}

	return errs
}


