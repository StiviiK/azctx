package azurecli

import (
	"encoding/json"
	"fmt"

	"github.com/StiviiK/azctx/log"
	"github.com/StiviiK/azctx/utils"
)

// Tenants returns all tenants
func (cli CLI) Tenants() []Tenant {
	return cli.tenants
}

// UpdateTenants fetches all available tenants from the azure management api and persists them to the azctxTenants.json file
func (cli *CLI) UpdateTenants() error {
	// Fetch all available tenants from the azure management api (https://management.azure.com/tenants?api-version=2020-01-01) using the azure cli
	args := []string{
		"rest",
		"--method", "get",
		"--url", "/tenants?api-version=2020-01-01",
		"--output", "json",
	}

	output, err := utils.ExecuteCommand(AZ_COMMAND, args...)
	if err != nil {
		return err
	}

	// Unmarshal the output
	var apiResponse struct {
		Tenants []Tenant `json:"value"`
	}
	err = json.Unmarshal([]byte(output), &apiResponse)
	if err != nil {
		return fmt.Errorf("failed to unmarshal az account list output: %s", err.Error())
	}

	// Copy the tenants to the CLI instance
	cli.tenants = make([]Tenant, len(apiResponse.Tenants))
	copy(cli.tenants, apiResponse.Tenants)

	// Persist the tenants to the json file
	err = cli.writeTenants()
	if err != nil {
		return fmt.Errorf("failed to write tenants to json file: %s", err.Error())
	}

	return nil
}

// MapTenanIdsToNames maps the tenant ids to the tenant names
func (cli *CLI) MapTenantIdsToNames() {
	if len(cli.tenants) > 0 {
		// Map the tenant ids to the tenant names
		tenantMap := make(map[string]string)
		for _, tenant := range cli.tenants {
			tenantMap[tenant.Id] = tenant.Name
		}

		// Rewrite the tenant ids to the tenant names
		for i, subscription := range cli.profile.Subscriptions {
			if tenantName, ok := tenantMap[subscription.Tenant]; ok {
				cli.profile.Subscriptions[i].Tenant = fmt.Sprintf("%s (%s)", tenantName, subscription.Tenant)
			}
		}
	} else {
		log.Info("If you want to fetch the tenant names, please authenticate the azure cli again using the wraper command: `azctx login`.")
	}
}
