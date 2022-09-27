package azurecli

import (
	"errors"

	"github.com/StiviiK/azctx/utils"
)

// Subscription represents a subscription in the AzureProfiles.json file
type Subscription struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Tenant    string `json:"tenantId"`
	IsDefault bool   `json:"isDefault"`
}

// subscriptionSorter is a custom sorter for subscriptions
type SubscriptionSorter []Subscription

func (a SubscriptionSorter) Len() int      { return len(a) }
func (a SubscriptionSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SubscriptionSorter) Less(i, j int) bool {
	subA, subB := a[i], a[j]

	if subA.Tenant == subB.Tenant {
		return subA.Name < subB.Name
	}

	return subA.Tenant > subB.Tenant
}

// SetSubscription sets the default subscription in the azure config file
func SetSubscription(subscription Subscription) error {
	// Execute az account set command
	_, err := utils.ExecuteCommand(Command, "account", "set", "--subscription", subscription.ID)
	if err != nil {
		return err
	}

	return nil
}

// ActiveSubscription returns the active subscription
func ActiveSubscription(profile Profile) (Subscription, error) {
	// select the subscription with the is default flag set to true
	for _, subscription := range profile.Subscriptions {
		if subscription.IsDefault {
			return subscription, nil
		}
	}

	return Subscription{}, errors.New("no active subscription found")
}

// SubscriptionNames returns the names of the given subscriptions
func SubscriptionNames(subscriptions []Subscription) []string {
	var subscriptionNames []string
	for _, subscription := range subscriptions {
		subscriptionNames = append(subscriptionNames, subscription.Name)
	}

	return subscriptionNames
}
