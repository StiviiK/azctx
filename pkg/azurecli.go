package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/StiviiK/azctx/azurecli"
	"github.com/StiviiK/azctx/log"
	"github.com/StiviiK/azctx/utils"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/spf13/afero"
)

const (
	AzureCLI_Command      = "az"
	AzureCLI_ConfigDirEnv = "AZURE_CONFIG_DIR"
	AzureCLI_ProfilesJSON = "azureProfile.json"
)

func EnsureAzureCLI() error {
	if !utils.IsCommandInstalled(AzureCLI_Command) {
		return errors.New("azure cli is not installed. please install it and try again. See here: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli")
	}

	return nil
}

// ReadAzureProfilesConfigFile reads the azure config file and returns the AzureProfilesConfig
func ReadAzureProfilesConfigFile(fs afero.Fs) (afero.File, error) {
	// Verify that the AZURE_CONFIG_DIR environment variable is set
	configDir := os.Getenv(AzureCLI_ConfigDirEnv)
	if configDir == "" {
		log.Warn("AZURE_CONFIG_DIR environment variable is not set. Using default config directory.")
		configDir = os.Getenv("HOME") + "/.azure"
	}

	// Verify that the config dir exists
	if !utils.FileExists(configDir) {
		return nil, fmt.Errorf("AZURE_CONFIG_DIR (%s) is not a valid directory. Please run `az configure` and try again", configDir)
	}

	// Verify that the azureProfile.json file exists
	configFilePath := fmt.Sprintf("%s/%s", configDir, AzureCLI_ProfilesJSON)
	if !utils.FileExists(configFilePath) {
		return nil, fmt.Errorf("%s is not a valid file. Please run `az configure` and try again", configFilePath)
	}

	// Open the azureProfile.json file
	configFile, err := fs.OpenFile(configFilePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid file: %s", configFilePath, err.Error())
	}

	return configFile, nil
}

// GetAzureProfileConfig returns the AzureProfilesConfig from the given azure config file
func GetAzureProfileConfig(profilesConfigFile afero.File) (azurecli.Profile, error) {
	// Read the azureProfile.json file
	configBytes, err := io.ReadAll(profilesConfigFile)
	if err != nil {
		return azurecli.Profile{}, fmt.Errorf("failed to read %s: %s", profilesConfigFile.Name(), err.Error())
	}

	// Remove UTF-8 BOM if present
	configBytes = utils.RemoveUTF8BOM(configBytes)

	// Unmarshal the config file
	var profilesConfigJSON azurecli.Profile
	err = json.Unmarshal(configBytes, &profilesConfigJSON)
	if err != nil {
		return azurecli.Profile{}, fmt.Errorf("failed to unmarshal %s: %s", profilesConfigFile.Name(), err.Error())
	}

	return profilesConfigJSON, nil
}

// ExecuteAzLogin executes the az login command
func ExecuteAzLogin(extraArgs []string) error {
	args := []string{"login", "--allow-no-subscriptions"}
	args = append(args, extraArgs...)
	err := utils.ExecuteCommandBare(AzureCLI_Command, os.Stdout, os.Stderr, args...)
	if err != nil {
		return err
	}

	return nil
}

// SetAzureSubscription sets the default subscription in the azure config file
func SetActiveSubscription(subscription azurecli.Subscription) error {
	// Execute az account set command
	_, err := utils.ExecuteCommand(AzureCLI_Command, "account", "set", "--subscription", subscription.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetActiveSubscription returns the active subscription
func GetActiveSubscription(profilesConfig azurecli.Profile) (azurecli.Subscription, error) {
	// select the subscription with the is default flag set to true
	for _, subscription := range profilesConfig.Subscriptions {
		if subscription.IsDefault {
			return subscription, nil
		}
	}

	return azurecli.Subscription{}, errors.New("no active subscription found")
}

// GetAzureSubscriptionByName returns the azure subscription with the given name
func GetAzureSubscriptionByName(profilesConfig azurecli.Profile, subscriptionName string) (azurecli.Subscription, bool) {
	// Find the subscription with the given name
	for _, subscription := range profilesConfig.Subscriptions {
		if strings.EqualFold(subscription.Name, subscriptionName) {
			return subscription, true
		}
	}

	return azurecli.Subscription{}, false
}

// TryFindSubscription fuzzy searches for the azure subscription in the given AzureProfilesConfig
func TryFindSubscription(profilesConfig azurecli.Profile, subscriptionName string) ([]azurecli.Subscription, error) {
	// Fuzzy search for the subscription name
	subscriptionNames := utils.StringSlice(azurecli.SubscriptionNames(profilesConfig.Subscriptions))
	results := fuzzy.FindNormalized(strings.ToLower(subscriptionName), subscriptionNames.ToLower())
	switch len(results) {
	case 0:
		// No results found
		return nil, fmt.Errorf("no azure subscription found for '%s'", subscriptionName)
	case 1:
		// One result found
		s, ok := GetAzureSubscriptionByName(profilesConfig, results[0])
		if !ok {
			return nil, fmt.Errorf("no azure subscription found for '%s'", subscriptionName)
		}
		return []azurecli.Subscription{s}, nil
	default:
		// Multiple results found
		subscriptions := make([]azurecli.Subscription, 0)
		for _, result := range results {
			s, ok := GetAzureSubscriptionByName(profilesConfig, result)
			if ok {
				subscriptions = append(subscriptions, s)
			}
		}

		return subscriptions, nil
	}
}
