package cmd

import (
	"github.com/StiviiK/azctx/azurecli"
	"github.com/StiviiK/azctx/utils"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Azure",
	Long:  `Login to Azure (wrapped around 'az login')`,
	Run:   utils.WrapCobraCommandHandler(testRunE),
}

func init() {

	rootCmd.AddCommand(loginCmd)
}

func testRunE(cmd *cobra.Command, args []string) error {
	// Ensure that the azure cli is installed
	err := azurecli.Ensure()
	if err != nil {
		return err
	}

	// Login to Azure
	err = azurecli.Login(args)
	if err != nil {
		return err
	}

	// Fetch all available tenants
	_, err = azurecli.FetchTenants()
	if err != nil {
		return err
	}

	return nil
}
