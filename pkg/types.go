package pkg

// AzureProfilesConfig represents the AzureProfiles.json file
type AzureProfilesConfig struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

// Subscription represents a subscription in the AzureProfiles.json file
type Subscription struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Tenant    string `json:"tenantId"`
	IsDefault bool   `json:"isDefault"`
}

// subscriptionSorter is a custom sorter for subscriptions
type subscriptionSorter []Subscription

func (a subscriptionSorter) Len() int           { return len(a) }
func (a subscriptionSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a subscriptionSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }
