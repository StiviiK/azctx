package utils

import (
	"bytes"
	"errors"
	"os/exec"
)

// IsCommandAvailable checks if a command is available
func IsCommandInstalled(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// ExecuteCommand executes a command and returns the output
func ExecuteCommand(name string, args ...string) (string, error) {
	// lookup the path of the command
	cmdPath, err := exec.LookPath(name)
	if err != nil {
		return "", err
	}

	// build the command
	stdoutBuffer := &bytes.Buffer{}
	errBuffer := &bytes.Buffer{}
	cmd := exec.Cmd{
		Path:   cmdPath,
		Args:   append([]string{name}, args...),
		Stdout: stdoutBuffer,
		Stderr: errBuffer,
	}

	// execute the command
	err = cmd.Run()
	if err != nil {
		return stdoutBuffer.String(), errors.New("Error executing command: " + errBuffer.String())
	}

	return stdoutBuffer.String(), nil
}
