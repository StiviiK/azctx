package azurecli

import (
	"encoding/json"
	"fmt"

	"github.com/StiviiK/azctx/utils"
)

// TenantListResponse represents the response of the azure management api
type TenantListResponse struct {
	Id   string `json:"tenantId"`
	Name string `json:"displayName"`
}

// FetchTenants fetches all available tenants from the azure management api
func FetchTenants() ([]TenantListResponse, error) {
	// Fetch all available tenants from the azure management api using the azure cli
	args := []string{
		"rest",
		"--method", "get",
		"--url", "/tenants?api-version=2020-01-01",
		"--output", "json",
	}

	output, err := utils.ExecuteCommand(Command, args...)
	if err != nil {
		return nil, err
	}

	// Unmarshal the output
	var apiResponse struct {
		Tenants []TenantListResponse `json:"value"`
	}
	err = json.Unmarshal([]byte(output), &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal az account list output: %s", err.Error())
	}

	return apiResponse.Tenants, nil
}

// FetchTenantNames rewrites the tenant ids to the tenant names
func FetchTenantNames(subscriptions []Subscription) []Subscription {
	// Fetch all available tenants
	tenants, err := FetchTenants()
	if err != nil {
		return subscriptions
	}

	// Map the tenant ids to the tenant names
	tenantMap := make(map[string]string)
	for _, tenant := range tenants {
		tenantMap[tenant.Id] = tenant.Name
	}

	// Rewrite the tenant ids to the tenant names
	for i, subscription := range subscriptions {
		if tenantName, ok := tenantMap[subscription.Tenant]; ok {
			subscriptions[i].Tenant = fmt.Sprintf("%s (%s)", tenantName, subscription.Tenant)
		}
	}

	return subscriptions
}
