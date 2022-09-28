package azurecli

import (
	"errors"
	"fmt"
	"os"

	"github.com/StiviiK/azctx/log"
	"github.com/StiviiK/azctx/utils"
	"github.com/spf13/afero"
)

// New creates a new CLI instance
func New(fs afero.Fs) (CLI, error) {
	// Ensure that the azure cli is installed
	if !utils.IsCommandInstalled(command) {
		return CLI{}, errors.New("azure cli is not installed. please install it and try again. See here: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli")
	}

	// Create a new CLI instance
	cli := CLI{
		fs:      fs,
		tenants: make([]Tenant, 0),

		profile: nil,
	}

	// Read the azureProfile.json file
	err := cli.readProfile()
	if err != nil {
		return CLI{}, err
	}

	// Read the azctxTenants.json file
	err = cli.readTenants()
	if err != nil {
		return CLI{}, err
	}

	return cli, nil
}

// Login executes the az login command
func (cli CLI) Login(extraArgs []string) error {
	args := []string{"login"}
	args = append(args, extraArgs...)
	err := utils.ExecuteCommandBare(command, os.Stdout, os.Stderr, args...)
	if err != nil {
		return err
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
