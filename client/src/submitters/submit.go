package submitters

import (
	"os"
	"fmt"
	"errors"
	"os/exec"
)

type Submitter func(string, string, string, map[string] string) error

func GetSubmitter() Submitter {
	var s Submitter = SubmitDocker
	if !isOnPath("docker") {
		fmt.Println("docker was not found on the path, using cURL instead")
		return SubmitOverHttp
	}
	return s
}

func isOnPath(cmd string) bool {
	_, err := exec.LookPath(cmd);
	return err == nil
}

func SubmitDocker(imageTag string, masterUrl string, swarmUri string, env map[string]string) error {
	suabCmd := buildSuabCmd(imageTag, masterUrl)

	cmd := exec.Command("docker", "run")
	cmd.Args = append(cmd.Args, "--entrypoint=/bin/bash")
	cmd.Args = appendEnv(cmd.Args, env)
	cmd.Args = append(cmd.Args, imageTag, "-c", suabCmd)

	if swarmUri != "magic:0" {
		cmd.Env = []string{"DOCKER_HOST=" + swarmUri}
	}
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
func appendEnv(args []string, env map[string]string) []string {
	for key, value := range env {
		args = append(args, "--env" ,key +"="+value)
	}
	return args
}

func buildSuabCmd(imageTag string, masterUrl string) string {
	exportBuildId := "export SUAB_BUILD_ID=`hostname`"
	baseUrl := masterUrl+ "/build/$SUAB_BUILD_ID"
	logFile := "/tmp/run-output"

	echoBuildId := "echo \"BuildId: $SUAB_BUILD_ID\""
	tellMasterThatABuildHasStarted := "curl --data '{\"image\": \"" +imageTag+ "\"}' " +baseUrl // TODO: Find what flags to use to only output things if it fails
	checkoutCode := "checkout-code.sh 2>&1 | tee " + logFile
	run := "run.sh 2>&1 | tee --apend " + logFile
	uploadLogs := "curl --data @" +logFile+ " " +baseUrl+ "/logs"
	uploadArtifacts := "test -d /artifacts && find /artifacts -type f -exec curl -X POST --data-binary @{} " +baseUrl+ "{} \\;"
	exitWithTheExitCodeFromRun := "exit 0" // TODO: Read the real exit code and exit with this.

	suabCmd := exportBuildId
	suabCmd += " && " + echoBuildId + " ; " + tellMasterThatABuildHasStarted + " ; " + checkoutCode
	suabCmd += " && " + run
	suabCmd += " && " + uploadLogs + " ; " + uploadArtifacts // TODO: The logs should be streamed to the server, not uploaded once it's all done
	suabCmd += " && " + exitWithTheExitCodeFromRun // TODO: If the uploads fail, which exit code do we want to use then? If run.sh was ok, then the uploads? Otherwise, let the run.sh trump?
	return suabCmd
}

func SubmitOverHttp(imageTag string, masterUrl string, swarmUri string, env map[string]string) error {
	// TODO
	return errors.New("Not implemented yet. Please install Docker and make sure it's on the path")
}
