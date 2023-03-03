package azurecli

import "github.com/spf13/afero"

// CLI represents the azure cli
type CLI struct {
	fs      afero.Fs
	profile Profile
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

// Subscription represents a subscription in the AzureProfiles.json file
type Subscription struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
	User  struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"user"`
	IsDefault    bool   `json:"isDefault"`
	Tenant       string `json:"tenantId"`
	TenantName   string `` // This is a duplicate of Tenant, used only for display purposes
	Environment  string `json:"environmentName"`
	HomeTenantId string `json:"homeTenantId"`
	ManagedBy    []struct {
		Id string `json:"tenantId"`
	} `json:"managedByTenants"`
}

// SubscriptionSlice is a custom sorter for subscriptions
type SubscriptionSlice []Subscription

type Slice[T any] []T
