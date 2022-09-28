package azurecli

import (
	"encoding/json"
	"fmt"

	"github.com/StiviiK/azctx/utils"
)

// UpdateTenants fetches all available tenants from the azure management api and persists them to the azctxTenants.json file
func (cli *CLI) UpdateTenants() error {
	// Fetch all available tenants from the azure management api using the azure cli
	args := []string{
		"rest",
		"--method", "get",
		"--url", "/tenants?api-version=2020-01-01",
		"--output", "json",
	}

	output, err := utils.ExecuteCommand(command, args...)
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
