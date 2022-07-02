package main

import (
	"github.com/StiviiK/azctx/log"
	"github.com/StiviiK/azctx/pkg"
	"github.com/StiviiK/azctx/utils"
	"github.com/spf13/afero"
)

func main() {
	// Initialize the file system abstraction
	fs := afero.NewOsFs()

	// Ensure that the azure cli is installed
	err := utils.EnsureAzureCLI()
	if err != nil {
		log.Error(err.Error())
		return
	}

	// Read the azure profiles config file
	profilesConfig, err := utils.ReadAzureProfilesConfig(fs)
	if err != nil {
		log.Error(err.Error())
		return
	}

	// Get the available subscriptions
	subscriptions, err := pkg.AvailableSubscriptions(profilesConfig)
	if err != nil {
		log.Error(err.Error())
		return
	}

	// Ask the user to select a subscription
	prompt := pkg.BuildPrompt(subscriptions)
	idx, _, err := prompt.Run()
	if err != nil {
		log.Error(err.Error())
		return
	}

	// Set the selected subscription as the default
	log.Info("Setting active subscription to %s (%s)", subscriptions[idx].Name, subscriptions[idx].ID)
	pkg.SetActiveSubscription(subscriptions[idx])
}
