package cmd

import (
	"errors"
	"os"

	"github.com/StiviiK/azctx/log"
	"github.com/StiviiK/azctx/pkg"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "azctx [- / NAME]",
	Short: "azctx is a CLI tool for managing azure cli subscriptions",
	Long: `azctx is a CLI tool for managing azure cli subscriptions.
	It is a helper for the azure cli and provides a simple interface for managing subscriptions.
	Pass a subscription name to select a specific subscription.
	Pass - to switch to the previous subscription.`,
	SilenceUsage: true,
	RunE:         cobraRunE,
	ValidArgs:    []string{"-", "NAME"},
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
		return selectSubscriptionByName(config, args[0])
	default:
		return interactivelySelectSubscription(config)
	}
}

func interactivelySelectSubscription(profilesConfig pkg.AzureProfilesConfig) error {
	// Ask the user to select a subscription
	subscriptions := profilesConfig.Subscriptions
	prompt := pkg.BuildPrompt(subscriptions)
	idx, _, err := prompt.Run()
	if err != nil {
		return nil
	}

	// Set the selected subscription as the default
	log.Info("Setting active subscription to %s (%s)", subscriptions[idx].Name, subscriptions[idx].ID)
	pkg.SetActiveSubscription(subscriptions[idx])

	return nil
}

func selectSubscriptionByName(profilesConfig pkg.AzureProfilesConfig, name string) error {
	subscriptions, err := pkg.TryFindAzureSubscription(profilesConfig, name)
	if err != nil {
		return err
	}

	// Check if we found more than one subscription
	var subscription pkg.Subscription
	if len(subscriptions) > 1 {
		prompt := pkg.BuildPrompt(subscriptions)
		idx, _, err := prompt.Run()
		if err != nil {
			return nil
		}

		subscription = subscriptions[idx]
	} else if len(subscriptions) == 1 {
		subscription = subscriptions[0]
	} else {
		return errors.New("no subscriptions found")
	}

	// Set the selected subscription as the default
	log.Info("Setting active subscription to %s (%s)", subscription.Name, subscription.ID)
	pkg.SetActiveSubscription(subscription)

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
	rootCmd.PersistentFlags().BoolP("current", "c", false, "Display the current active subscription")
	rootCmd.PersistentFlags().BoolP("refresh", "r", false, "Re-Authenticate and refresh the subscriptions")
}
