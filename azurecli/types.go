package azurecli

import "github.com/spf13/afero"

// CLI represents the azure cli
type CLI struct {
	fs      afero.Fs
	profile *Profile
	tenants []Tenant
}

// Profile represents the AzureProfiles.json file
type Profile struct {
	Subscriptions  []Subscription `json:"subscriptions"`
	InstallationId string         `json:"installationId"`
}

// Tenant represents the response of the azure management api
type Tenant struct {
	Id   string `json:"tenantId"`
	Name string `json:"displayName"`
}

// Subscriptions
type subscriptionUser struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type managingTenant struct {
	ID string `json:"tenantId"`
}

// Subscription represents a subscription in the AzureProfiles.json file
type Subscription struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	State        string           `json:"state"`
	User         subscriptionUser `json:"user"`
	IsDefault    bool             `json:"isDefault"`
	Tenant       string           `json:"tenantId"`
	Environment  string           `json:"environmentName"`
	HomeTenantId string           `json:"homeTenantId"`
	ManagedBy    []managingTenant `json:"managedByTenants"`
}

// SubscriptionSlice is a custom sorter for subscriptions
type SubscriptionSlice []Subscription
