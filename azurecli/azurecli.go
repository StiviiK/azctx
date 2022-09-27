package azurecli

import (
	"errors"
	"os"

	"github.com/StiviiK/azctx/utils"
)

const (
	Command      = "az"
	ConfigDirEnv = "AZURE_CONFIG_DIR"
	ProfilesJson = "azureProfile.json"
)

// Profile represents the AzureProfiles.json file
type Profile struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

// Ensure ensures that the azure cli is installed
func Ensure() error {
	if !utils.IsCommandInstalled(Command) {
		return errors.New("azure cli is not installed. please install it and try again. See here: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli")
	}

	return nil
}

// Login executes the az login command
func Login(extraArgs []string) error {
	args := []string{"login", "--allow-no-subscriptions"}
	args = append(args, extraArgs...)
	err := utils.ExecuteCommandBare(Command, os.Stdout, os.Stderr, args...)
	if err != nil {
		return err
	}

	return nil
}
