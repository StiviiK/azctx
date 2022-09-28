package azurecli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/StiviiK/azctx/utils"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

// Subscriptions returns all subscriptions
func (cli CLI) Subscriptions() []Subscription {
	return cli.profile.Subscriptions
}

// SubscriptionNames returns the names of the given subscriptions
func (cli CLI) SubscriptionNames() []string {
	return SubscriptionSlice(cli.profile.Subscriptions).SubscriptionNames()
}

// SetSubscription sets the default subscription in the azure config file
func (cli *CLI) SetSubscription(subscription Subscription) error {
	// Execute az account set command
	_, err := utils.ExecuteCommand(command, "account", "set", "--subscription", subscription.ID)
	if err != nil {
		return err
	}

	return nil
}

// ActiveSubscription returns the active subscription
func (cli CLI) ActiveSubscription() (Subscription, error) {
	// select the subscription with the is default flag set to true
	for _, subscription := range cli.profile.Subscriptions {
		if subscription.IsDefault {
			return subscription, nil
		}
	}

	return Subscription{}, errors.New("no active subscription found")
}

// GetSubscriptionByName returns the azure subscription with the given name
func (cli CLI) GetSubscriptionByName(subscriptionName string) (Subscription, bool) {
	// Find the subscription with the given name
	for _, subscription := range cli.profile.Subscriptions {
		if strings.EqualFold(subscription.Name, subscriptionName) {
			return subscription, true
		}
	}

	return Subscription{}, false
}

// TryFindSubscription fuzzy searches for the azure subscription in the given AzureProfilesConfig
func (cli CLI) TryFindSubscription(subscriptionName string) ([]Subscription, error) {
	// Fuzzy search for the subscription name
	subscriptionNames := utils.StringSlice(cli.SubscriptionNames())
	results := fuzzy.FindNormalized(strings.ToLower(subscriptionName), subscriptionNames.ToLower())

	switch len(results) {
	case 0:
		// No results found
		return nil, fmt.Errorf("no azure subscription found for '%s'", subscriptionName)
	case 1:
		// One result found
		s, ok := cli.GetSubscriptionByName(results[0])
		if !ok {
			return nil, fmt.Errorf("no azure subscription found for '%s'", subscriptionName)
		}
		return []Subscription{s}, nil
	default:
		// Multiple results found
		subscriptions := make([]Subscription, 0)
		for _, result := range results {
			s, ok := cli.GetSubscriptionByName(result)
			if ok {
				subscriptions = append(subscriptions, s)
			}
		}

		return subscriptions, nil
	}
}

func (a SubscriptionSlice) Len() int      { return len(a) }
func (a SubscriptionSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SubscriptionSlice) Less(i, j int) bool {
	subA, subB := a[i], a[j]

	if subA.Tenant == subB.Tenant {
		return subA.Name < subB.Name
	}

	return subA.Tenant > subB.Tenant
}

// SubscriptionNames returns the names of the given subscriptions
func (subscriptionSlice SubscriptionSlice) SubscriptionNames() []string {
	var subscriptionNames []string
	for _, subscription := range subscriptionSlice {
		subscriptionNames = append(subscriptionNames, subscription.Name)
	}

	return subscriptionNames
}
