package main

import (
	"fmt"
	"submitters"
	"config"
	"os"
	"strings"
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

	conf, err := config.ReadAndParseEffectiveConf(configFilePath);
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	if errs := validate(conf, configFilePath); len(errs) > 0  {
		fmt.Printf("Invalid config.\n  %s\n", strings.Join(errs, "\n  "))
		fmt.Println("\nUse the -h flag for more help")
		os.Exit(2)
	}

	return conf
}

func validate(conf *config.Config, configFile string) []string {
	errs := make([]string, 0)

	if len(conf.DockerImageTag) == 0 {
		errs = append(errs, "You must specify a docker image tag, either via the -d flag, or via the \"dockerImageTag\": \"some-tag\" in " + configFile)
	}

	if conf.MasterUrl == "" {
		errs = append(errs, "You must specify a valid master url, either via the -m flag, or via the \"masterUrl\": \"http://example.com:8080\" in " + configFile)
	}

	if conf.SwarmUri == "" {
		errs = append(errs, "You must specify a valid docker swarm uri, either via the -s flag, or via the \"swarmUri\": \"example.com:4000\" in " + configFile)
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

