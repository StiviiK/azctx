package utils

import (
	"fmt"
	"os/exec"
)

// IsCommandAvailable checks if a command is available
func IsCommandInstalled(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// ExecuteCommand executes a command and returns the output
func ExecuteCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.Output()

	fmt.Printf("%s\n", output)

	if err != nil {
		return "", err
	}

	return string(output), nil
}
