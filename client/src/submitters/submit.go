package submitters

import (
	"os"
	"fmt"
	"errors"
	"os/exec"
	"syscall"
	"strings"
	"regexp"
	"bytes"
)

type Submitter func(string, string, string, string, map[string] string) (int, error)

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

func SubmitDocker(suabShellScript string, imageTag string, masterUrl string, swarmUri string, env map[string]string) (int, error) {
	suabCmd := collapseToOneLine(suabShellScript)

	cmd := exec.Command("docker", "run")
	cmd.Args = append(cmd.Args, "--entrypoint=/bin/bash")
	cmd.Args = appendEnv(cmd.Args, env)
	cmd.Args = append(cmd.Args, imageTag, "-c", suabCmd)

	if len(swarmUri) > 0 {
		cmd.Env = []string{"DOCKER_HOST=" + swarmUri}
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Submitting image %s to %s. Posting the results to %s\n", imageTag, swarmUri, masterUrl)
	err := cmd.Run()

	if err == nil {
		// All was good!
		return 0, nil
	} else {
		// Something went wrong when executing the docker command. Let's try to find out what
		return tryToFindExitCode(err)
	}
}

func collapseToOneLine(script string) string {
	commentsRegex := regexp.MustCompile("#.*")
	var buffer bytes.Buffer

	lines := strings.Split(script, "\n")
	for i, line := range lines {
		lineWithoutComments := commentsRegex.ReplaceAllString(line, "")
		trimmedLine := strings.TrimSpace(lineWithoutComments)

		if len(trimmedLine) > 0 {
			buffer.WriteString(line)

			if i < len(lines) -1 && line != "(" {
				buffer.WriteString(";")
			}
		}
	}
	return buffer.String()
}

func appendEnv(args []string, env map[string]string) []string {
	for key, value := range env {
		args = append(args, "--env" ,key +"="+value)
	}
	return args
}

func tryToFindExitCode(err error) (int, error) {
	// Is the error an exec.ExitError?
	if exiterr, ok := err.(*exec.ExitError); ok {
		// Yes! This implies that the command exited with a non-zero exit code

		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			// We found the exit code!

			return status.ExitStatus(), nil
		} else {
			// The error was caused by the command exiting with a non-zero code, but we don't
			// know how to find this code..
			return 0, errors.New("The command exited with a non-zero code but we can't find out which one")
		}
	} else {
		// err was not an exec.ExitError, this might be a problem setting up the command

		s := fmt.Sprintf("Failed to run the command. %s\n", err)
		return 0, errors.New(s)
	}
}

func SubmitOverHttp(suabCmd string, imageTag string, masterUrl string, swarmUri string, env map[string]string) (int, error) {
	// TODO
	return 0, errors.New("Not implemented yet. Please install Docker and make sure it's on the path")
}
