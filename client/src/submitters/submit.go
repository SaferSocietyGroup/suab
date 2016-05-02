package submitters

import (
	"os"
	"fmt"
	"errors"
	"strings"
	"os/exec"
)

type Submitter func(string, string, string) error

func GetSubmitter() Submitter {
	var s Submitter = SumbitDocker
	if !isOnPath("docker") {
		fmt.Println("docker was not found on the path, using cURL instead")
		return SubmibOverHttp
	}
	return s
}

func isOnPath(cmd string) bool {
	_, err := exec.LookPath(cmd);
	return err == nil
}

func SumbitDocker(imageTag string, masterUrl string, swarmUri string) error {
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

func SubmibOverHttp(imageTag string, masterUrl string, swarmUri string) error {
	// TODO
	return nil
}
