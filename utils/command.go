package utils

import "os/exec"

func IsCommandInstalled(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
