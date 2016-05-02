package main

import (
	"fmt"
	"os/exec"
	"errors"
	"strings"
	"os"
)

type submitter func(string, string, string) error

func sumbitDocker(imageTag string, masterUrl string, swarmUri string) error {
	suabCmd := buildSuabCmd(imageTag, masterUrl)

	cmd := exec.Command("docker", "run", "--entrypoint=/bin/bash", /*"--rm",*/ imageTag, "-c", suabCmd)
	cmd.Env = []string{"DOCKER_HOST=" + swarmUri}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Submitting image %s to %s. Posting the results to %s\n", imageTag, swarmUri, masterUrl)
	err := cmd.Run()
	if err != nil {
		s := fmt.Sprintf("Could not submit over docker. %s\n", err)
		return errors.New(s)
	}

	return nil
}

func buildSuabCmd(imageTag string, masterUrl string) string {
	buildId := "`hostname`"
	baseUrl := masterUrl+ "/build/" + buildId
	logFile := "/tmp/run-output" 

	tellMasterThatABuildHasStarted := "curl --data '{\"hellur\": \"knivur\"}' " +baseUrl
	uploadLogs := "curl --data @" +logFile+ " " +baseUrl+ "/logs"
	uploadArtifacts := "test -d /artifacts && find /artifacts -type f -exec curl -X POST --data-binary @{} " +baseUrl+ "{} \\;"

	suabCmd := strings.Join([]string{
		"echo \"BuildId: " + buildId + "\"",
		tellMasterThatABuildHasStarted + " ; " +
		"checkout-code.sh", // TODO: We need arguments to the checkout-code.sh script. The hash to checkout e.g.
		"run.sh 2>&1 | tee " + logFile,
		uploadLogs + " ; " + // TODO: The logs should be streamed to the server, not uploaded once it's all done
		uploadArtifacts,
	}, " && ")
	return suabCmd
}

func submibOverHttp(imageTag string, masterUrl string, swarmUri string) error {
	// TODO
	return nil
}

func main() {
	// I want to run suab with no arguments to build the project I'm in
	// The following three variables should come from a config file 
	// and be overridable by flags

	dockerImageTag := os.Args[1]//"192.168.10.78:5000/apa"
	masterUrl := "http://192.168.10.78:8080"
	swarmUri := "192.168.10.78:4000"

	s := getSubmitter()
	err := s(dockerImageTag, masterUrl, swarmUri)
	if err != nil {
		fmt.Printf("Submission failed. %s\n", err)
	} else {
		fmt.Println("Successfully shut up and built")
	}
}

func getSubmitter() submitter {
	var s submitter = sumbitDocker
	if !isOnPath("docker") {
		fmt.Println("docker was not found on the path, using cURL instead")
		return submibOverHttp
	}
	return s
}

func isOnPath(cmd string) bool {
	_, err := exec.LookPath(cmd);
	return err == nil
}
