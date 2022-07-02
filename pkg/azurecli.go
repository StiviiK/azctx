package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/StiviiK/azctx/utils"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/spf13/afero"
)

type AzureProfilesConfig struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

type Subscription struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Tenant string `json:"tenantId"`
}

const (
	AzureCLI_Command      = "az"
	AzureCLI_ConfigDirEnv = "AZURE_CONFIG_DIR"
	AzureCLI_ProfilesJSON = "azureProfile.json"
)

func EnsureAzureCLI() error {
	if !utils.IsCommandInstalled(AzureCLI_Command) {
		return errors.New("azure cli is not installed. Please install it and try again.\nhttps://docs.microsoft.com/en-us/cli/azure/install-azure-cli")
	}

	return nil
}

// ReadAzureProfilesConfigFile reads the azure config file and returns the AzureProfilesConfig
func ReadAzureProfilesConfigFile(fs afero.Fs) (afero.File, error) {
	// Verify that the AZURE_CONFIG_DIR environment variable is set
	configDir := os.Getenv(AzureCLI_ConfigDirEnv)
	if configDir == "" {
		return nil, errors.New("AZURE_CONFIG_DIR is not set. Please create it and try again")
	}

	// Verify that the config dir exists
	if !utils.FileExists(configDir) {
		return nil, errors.New("AZURE_CONFIG_DIR is not a valid directory. Please run `az configure` and try again")
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
func GetAzureProfileConfig(profilesConfigFile afero.File) (AzureProfilesConfig, error) {
	// Read the azureProfile.json file
	configBytes, err := ioutil.ReadAll(profilesConfigFile)
	if err != nil {
		return AzureProfilesConfig{}, fmt.Errorf("failed to read %s: %s", profilesConfigFile.Name(), err.Error())
	}

	// Remove UTF-8 BOM if present
	configBytes = utils.RemoveUTF8BOM(configBytes)

	// Unmarshal the config file
	var profilesConfigJSON AzureProfilesConfig
	err = json.Unmarshal(configBytes, &profilesConfigJSON)
	if err != nil {
		return AzureProfilesConfig{}, fmt.Errorf("failed to unmarshal %s: %s", profilesConfigFile.Name(), err.Error())
	}

	return profilesConfigJSON, nil
}

// SetAzureSubscription sets the default subscription in the azure config file
func SetActiveSubscription(subscription Subscription) error {
	// Execute az account set command
	_, err := utils.ExecuteCommand("az", "account", "set", "--subscription", subscription.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetAzureSubscriptionByName returns the azure subscription with the given name
func GetAzureSubscriptionByName(profilesConfig AzureProfilesConfig, subscriptionName string) Subscription {
	// Find the subscription with the given name
	for _, subscription := range profilesConfig.Subscriptions {
		if subscription.Name == subscriptionName {
			return subscription
		}
	}

	return Subscription{}
}

// TryFindAzureSubscription fuzzy searches for the azure subscription in the given AzureProfilesConfig
func TryFindAzureSubscription(profilesConfig AzureProfilesConfig, subscriptionName string) ([]Subscription, error) {
	// Fuzzy search for the subscription name
	results := fuzzy.FindNormalized(strings.ToLower(subscriptionName), utils.LowercaseStrings(GetAzureSubscriptionNames(profilesConfig.Subscriptions)))
	switch len(results) {
	case 0:
		// No results found
		return nil, fmt.Errorf("no azure subscription found for %s", subscriptionName)
	case 1:
		// One result found
		return []Subscription{GetAzureSubscriptionByName(profilesConfig, results[0])}, nil
	default:
		// Multiple results found
		subscriptions := make([]Subscription, len(results))
		for i, result := range results {
			s := GetAzureSubscriptionByName(profilesConfig, result)
			if (s != Subscription{}) {
				subscriptions[i] = s
			}
		}

		return subscriptions, nil
	}
}

// GetAzureSubscriptionNames returns the names of the given subscriptions as a slice of strings
func GetAzureSubscriptionNames(subscriptions []Subscription) []string {
	var subscriptionNames []string
	for _, subscription := range subscriptions {
		subscriptionNames = append(subscriptionNames, subscription.Name)
	}

	return subscriptionNames
}
