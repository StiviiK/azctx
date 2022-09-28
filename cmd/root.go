package cmd

import (
	"errors"
	"os"

	"github.com/StiviiK/azctx/azurecli"
	"github.com/StiviiK/azctx/log"
	"github.com/StiviiK/azctx/prompt"
	"github.com/StiviiK/azctx/updates"
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
	Run:          utils.WrapCobraCommandHandler(rootRunE),
	ValidArgs:    []string{"-", "NAME"},
}

func init() {
	rootCmd.AddCommand(
		extension.NewVersionCobraCmd(
			extension.WithUpgradeNotice(owner, repo),
		),
	)

	rootCmd.Flags().BoolP("current", "c", false, "Display the current active subscription")
	rootCmd.Flags().BoolP("refresh", "r", false, `Re-Authenticate and refresh the subscriptions. 
	Deprecated. Please use azctx login instead.`)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func rootRunE(cmd *cobra.Command, args []string) error {
	// Initialize the CLI
	cli, err := azurecli.New(afero.NewOsFs())
	if err != nil {
		return err
	}

	// Try to map the tenant ids to names
	cli.MapTenantIdsToNames()

	// check the flags
	switch {
	case cmd.Flags().Changed("current"):
		return getActiveSubscription(cli)
	case cmd.Flags().Changed("refresh"):
		log.Warn("Deprecated flag --refresh/-r used. Please use `azctx login` instead.")
		return refreshSubscriptions(cli, args)
	case len(args) == 1:
		// check if the user passed -
		if args[0] == "-" {
			return errors.New("not implemented")
		}

		err = selectSubscriptionByName(cli, args[0])
		if err != nil {
			return err
		}
	default:
		err = interactivelySelectSubscription(cli)
		if err != nil {
			return err
		}
	}

	return updates.Check(owner, repo)
}

func interactivelySelectSubscription(cli azurecli.CLI) error {
	// Ask the user to select a subscription
	prompt := prompt.BuildPrompt(cli.Subscriptions())
	idx, _, err := prompt.Run()
	if err != nil {
		return nil
	}

	// Set the selected subscription as the default
	subscriptions := cli.Subscriptions()
	log.Info("Setting active subscription to %s (%s)", subscriptions[idx].Name, subscriptions[idx].ID)
	err = cli.SetSubscription(subscriptions[idx])
	if err != nil {
		return err
	}

	return nil
}

func selectSubscriptionByName(cli azurecli.CLI, name string) error {
	log.Info(name)
	subscriptions, err := cli.TryFindSubscription(name)
	if err != nil {
		return err
	}

	// Check if we found more than one subscription
	var subscription azurecli.Subscription
	switch length := len(subscriptions); {
	case length > 1:
		prompt := prompt.BuildPrompt(subscriptions)
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
	err = cli.SetSubscription(subscription)
	if err != nil {
		return err
	}

	return nil
}

func getActiveSubscription(cli azurecli.CLI) error {
	// Try to get the active subscription
	subscription, err := cli.ActiveSubscription()
	if err != nil {
		return err
	}

	// Print the active subscription
	log.Info("Active subscription: %s (%s)", subscription.Name, subscription.ID)

	return nil
}
