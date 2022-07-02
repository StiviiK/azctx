package pkg

import (
	"encoding/json"
	"io/ioutil"

	"github.com/StiviiK/azctx/utils"
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

func AvailableSubscriptions(profilesConfig afero.File) ([]Subscription, error) {
	// Read the azure profiles config file
	var profilesConfigJSON AzureProfilesConfig
	data, err := ioutil.ReadAll(profilesConfig)
	if err != nil {
		return nil, err
	}

	// Remove UTF-8 BOM if present
	data = utils.RemoveUTF8BOM(data)

	// Unmarshal the JSON
	err = json.Unmarshal(data, &profilesConfigJSON)
	if err != nil {
		return nil, err
	}

	return profilesConfigJSON.Subscriptions, nil
}

// GetAzureSubscriptionNames returns the names of the subscriptions in the azure config file as a slice of strings
func GetAzureSubscriptionNames(subscriptions []Subscription) []string {
	var subscriptionNames []string
	for _, subscription := range subscriptions {
		subscriptionNames = append(subscriptionNames, subscription.Name)
	}

	return subscriptionNames
}

// SetAzureSubscription sets the default subscription in the azure config file
func SetActiveSubscription(s Subscription) error {
	// Execute az account set command
	_, err := utils.ExecuteCommand("azf", "account", "set", "--subscription", s.ID)
	if err != nil {
		return err
	}

	return nil
}
