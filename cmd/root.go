package cmd

import (
	"errors"
	"os"

	"github.com/StiviiK/azctx/log"
	"github.com/StiviiK/azctx/pkg"
	"github.com/StiviiK/azctx/utils"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.szostok.io/version/extension"
)

const (
	// Respository Owner
	owner = "stiviik"

	// Repository Name
	repo = "azctx"
)

var rootCmd = &cobra.Command{
	Use:   "azctx [- / NAME]",
	Short: "azctx is a CLI tool for managing azure cli subscriptions",
	Long: `azctx is a CLI tool for managing azure cli subscriptions.
	It is a helper for the azure cli and provides a simple interface for managing subscriptions.
	Pass a subscription name to select a specific subscription.
	Pass - to switch to the previous subscription.`,
	SilenceUsage: true,
	Run:          utils.WrapCobraCommandHandler(cobraRunE),
	ValidArgs:    []string{"-", "NAME"},
}

func init() {
	rootCmd.AddCommand(
		extension.NewVersionCobraCmd(
			extension.WithUpgradeNotice(owner, repo),
		),
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func cobraRunE(cmd *cobra.Command, args []string) error {
	// Initialize the file system abstraction
	fs := afero.NewOsFs()

	// Ensure that the azure cli is installed
	err := pkg.EnsureAzureCLI()
	if err != nil {
		return err
	}

	// Read the azure profiles config file
	profilesConfig, err := pkg.ReadAzureProfilesConfigFile(fs)
	if err != nil {
		return err
	}

	// Paarse the azure profile config
	config, err := pkg.GetAzureProfileConfig(profilesConfig)
	if err != nil {
		return err
	}

	// check if we should fetch the tenant names
	fetchTenantNames, err := cmd.Flags().GetBool("fetch-tenant-names")
	if err != nil {
		return err
	}

	// check the flags
	switch {
	case cmd.Flags().Changed("current"):
		return getActiveSubscription(config)
	case cmd.Flags().Changed("refresh"):
		return errors.New("not implemented")
	case len(args) == 1:
		// check if the user passed -
		if args[0] == "-" {
			return errors.New("not implemented")
		}

		err = selectSubscriptionByName(config, args[0], fetchTenantNames)
		if err != nil {
			return err
		}
	default:
		err = interactivelySelectSubscription(config, fetchTenantNames)
		if err != nil {
			return err
		}
	}

	return pkg.CheckForUpdates(owner, repo)
}

func interactivelySelectSubscription(profilesConfig pkg.AzureProfilesConfig, fetchTenantNames bool) error {
	// Ask the user to select a subscription
	subscriptions := profilesConfig.Subscriptions

	// Check if we should fetch the tenant names
	if fetchTenantNames {
		// Fetch the tenant names
		subscriptions = pkg.FetchTenantNames(subscriptions)
	}

	prompt := pkg.BuildPrompt(subscriptions)
	idx, _, err := prompt.Run()
	if err != nil {
		return nil
	}

	// Set the selected subscription as the default
	log.Info("Setting active subscription to %s (%s)", subscriptions[idx].Name, subscriptions[idx].ID)
	err = pkg.SetActiveSubscription(subscriptions[idx])
	if err != nil {
		return err
	}

	return nil
}

func selectSubscriptionByName(profilesConfig pkg.AzureProfilesConfig, name string, fetchTenantNames bool) error {
	subscriptions, err := pkg.TryFindAzureSubscription(profilesConfig, name)
	if err != nil {
		return err
	}

	// Check if we found more than one subscription
	var subscription pkg.Subscription
	switch length := len(subscriptions); {
	case length > 1:
		// Check if we should fetch the tenant names
		if fetchTenantNames {
			subscriptions = pkg.FetchTenantNames(subscriptions)
		}

		prompt := pkg.BuildPrompt(subscriptions)
		idx, _, err := prompt.Run()
		if err != nil {
			return nil
		}

		subscription = subscriptions[idx]
	case length == 1:
		subscription = subscriptions[0]
	default:
		return errors.New("no subscription found")
	}

	// Set the selected subscription as the default
	log.Info("Setting active subscription to %s (%s)", subscription.Name, subscription.ID)
	err = pkg.SetActiveSubscription(subscription)
	if err != nil {
		return err
	}

	return nil
}

func getActiveSubscription(profilesConfig pkg.AzureProfilesConfig) error {
	// Try to get the active subscription
	subscription, err := pkg.GetActiveSubscription(profilesConfig)
	if err != nil {
		return err
	}

	// Print the active subscription
	log.Info("Active subscription: %s (%s)", subscription.Name, subscription.ID)

	return nil
}

func init() {
	rootCmd.Flags().BoolP("current", "c", false, "Display the current active subscription")
	rootCmd.Flags().BoolP("refresh", "r", false, "Re-Authenticate and refresh the subscriptions")
	rootCmd.Flags().BoolP("fetch-tenant-names", "t", false, "Fetch the tenant names for the subscriptions")
}
