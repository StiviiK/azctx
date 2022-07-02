package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/afero"
)

type AzureProfilesConfig struct {
	Subscriptions []AzureProfile `json:"subscriptions"`
}

type AzureProfile struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"isDefault"`
	TenantId  string `json:"tenantId"`
}

const (
	AzureCLI_Command      = "az"
	AzureCLI_ConfigDirEnv = "AZURE_CONFIG_DIR"
	AzureCLI_ProfilesJSON = "azureProfile.json"
)

func EnsureAzureCLI() error {
	if !IsCommandInstalled(AzureCLI_Command) {
		return errors.New("azure cli is not installed. Please install it and try again.\nhttps://docs.microsoft.com/en-us/cli/azure/install-azure-cli")
	}

	return nil
}

// ReadAzureProfilesConfig reads the azure config file and returns the AzureProfilesConfig
func ReadAzureProfilesConfig(fs afero.Fs) (afero.File, error) {
	// Verify that the AZURE_CONFIG_DIR environment variable is set
	configDir := os.Getenv(AzureCLI_ConfigDirEnv)
	if configDir == "" {
		return nil, errors.New("AZURE_CONFIG_DIR is not set. Please set it and try again")
	}

	// Verify that the config dir exists
	if !IsFileExists(configDir) {
		return nil, errors.New("AZURE_CONFIG_DIR is not a valid directory. Please run `az configure` and try again")
	}

	// Verify that the azureProfile.json file exists
	configFilePath := fmt.Sprintf("%s/%s", configDir, AzureCLI_ProfilesJSON)
	if !IsFileExists(configFilePath) {
		return nil, fmt.Errorf("%s is not a valid file. Please run `az configure` and try again", configFilePath)
	}

	// Open the azureProfile.json file
	configFile, err := fs.OpenFile(configFilePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid file: %s", configFilePath, err.Error())
	}

	return configFile, nil
}
