package main

import (
	"github.com/StiviiK/azctx/log"
	"github.com/StiviiK/azctx/pkg"
	"github.com/StiviiK/azctx/utils"
)

func main() {
	err := utils.EnsureAzureCLI()
	if err != nil {
		log.Error(err.Error())
	}

	prompt := pkg.BuildPrompt()
	prompt.Run()
}
