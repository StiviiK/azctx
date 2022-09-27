package utils

import (
	"bytes"
	"errors"
	"io"
	"os/exec"
)

// IsCommandAvailable checks if a command is available
func IsCommandInstalled(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// ExecuteCommandBare executes a command and writes the output to the given writers
func ExecuteCommandBare(name string, stdOut, stdErr io.Writer, args ...string) error {
	// lookup the path of the command
	cmdPath, err := exec.LookPath(name)
	if err != nil {
		return err
	}

	// build the command
	cmd := exec.Cmd{
		Path:   cmdPath,
		Args:   append([]string{name}, args...),
		Stdout: stdOut,
		Stderr: stdErr,
	}

	// execute the command
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// ExecuteCommand executes a command and returns the output
func ExecuteCommand(name string, args ...string) (string, error) {
	stdoutBuffer := &bytes.Buffer{}
	errBuffer := &bytes.Buffer{}
	err := ExecuteCommandBare(name, stdoutBuffer, errBuffer, args...)
	if err != nil {
		if errBuffer.Len() > 0 {
			return "", errors.New("Error executing command: " + errBuffer.String())
		} else {
			return "", err
		}
	}

	return stdoutBuffer.String(), nil
}
