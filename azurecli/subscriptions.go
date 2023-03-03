package azurecli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/StiviiK/azctx/utils"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

var (
	FilterTenantLevelAccount bool = true // Filter tenant level accounts with no available subscriptions, will be manipulated by the flag --filter-tenant-level in cmd/root.go#L47
)

// Subscriptions returns all subscriptions
func (cli CLI) Subscriptions() SubscriptionSlice {
	//return cli.profile.Subscriptions
	filter := func(s Subscription) bool {
		return !FilterTenantLevelAccount || !strings.EqualFold(s.Name, "N/A(tenant level account)")
	}
	return utils.Filter(cli.profile.Subscriptions, filter)
}

// SubscriptionNames returns the names of the given subscriptions
func (cli CLI) SubscriptionNames() utils.StringSlice {
	return cli.Subscriptions().Names()
}

// SetSubscription sets the default subscription in the azure config file
func (cli *CLI) SetSubscription(subscription Subscription) error {
	// Execute az account set command
	_, err := utils.ExecuteCommand(AZ_COMMAND, "account", "set", "--subscription", subscription.Id)
	if err != nil {
		return err
	}

	return nil
}

// ActiveSubscription returns the active subscription
func (cli CLI) ActiveSubscription() (Subscription, error) {
	// select the subscription with the is default flag set to true
	for _, subscription := range cli.profile.Subscriptions { // Do not use cli.Subscriptions() here, because we want to return the subscription even if it is a tenant level account
		if subscription.IsDefault {
			return subscription, nil
		}
	}

	return Subscription{}, errors.New("no active subscription found")
}

// GetSubscriptionByName returns the azure subscription with the given name
func (cli CLI) GetSubscriptionByName(subscriptionName string) (Subscription, bool) {
	// Find the subscription with the given name
	for _, subscription := range cli.Subscriptions() {
		if strings.EqualFold(subscription.Name, subscriptionName) {
			return subscription, true
		}
	}

	return Subscription{}, false
}

// TryFindSubscription fuzzy searches for the azure subscription in the given AzureProfilesConfig
func (cli CLI) TryFindSubscription(subscriptionName string) (SubscriptionSlice, error) {
	// Fuzzy search for the subscription name
	subscriptionNames := cli.SubscriptionNames()
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
			if s, ok := cli.GetSubscriptionByName(result); ok {
				subscriptions = append(subscriptions, s)
			}
		}

		return subscriptions, nil
	}
}

// implement sort.Interface for SubscriptionSlice
func (a SubscriptionSlice) Len() int      { return len(a) }
func (a SubscriptionSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SubscriptionSlice) Less(i, j int) bool {
	subA, subB := a[i], a[j]

	// Sort subscriptions by tenant name and then by subscription name
	if subA.TenantName == subB.TenantName {
		return subA.Name < subB.Name
	}

	return subA.TenantName < subB.TenantName
}

// Names returns the names of the given subscriptions
func (subscriptionSlice SubscriptionSlice) Names() []string {
	var subscriptionNames []string
	for _, subscription := range subscriptionSlice {
		subscriptionNames = append(subscriptionNames, subscription.Name)
	}

	return subscriptionNames
}
