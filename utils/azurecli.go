package utils

import (
	"errors"
)

const (
	AzureCLI_Command = "az"
)

func EnsureAzureCLI() error {
	if !IsCommandInstalled(AzureCLI_Command) {
		return errors.New("azure cli is not installed. Please install it and try again.\nhttps://docs.microsoft.com/en-us/cli/azure/install-azure-cli")
	}

	return nil
}
