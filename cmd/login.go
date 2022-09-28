package cmd

import (
	"github.com/StiviiK/azctx/azurecli"
	"github.com/StiviiK/azctx/utils"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Azure",
	Long: `Login to Azure (wrapped around 'az login')
	It will refresh the subscriptions and fetch the tenant names.
	All args after -- are directly passed to the 'az login' command.`,
	Run: utils.WrapCobraCommandHandler(loginRunE),
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func loginRunE(cmd *cobra.Command, args []string) error {
	// Ensure that the azure cli is installed
	cli, err := azurecli.New(afero.NewOsFs())
	if err != nil {
		return err
	}

	// Refresh the subscriptions
	err = refreshSubscriptions(cli, args)
	if err != nil {
		return err
	}

	return nil
}

func refreshSubscriptions(cli azurecli.CLI, extraArgs []string) error {
	// Try to refresh the subscriptions
	err := cli.Login(extraArgs)
	if err != nil {
		return err
	}

	// Fetch all available tenants
	err = cli.UpdateTenants()
	if err != nil {
		return err
	}

	return nil
}
