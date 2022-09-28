# `azctx`: Power tool for the Azure CLI

![Proudly written in Go](https://img.shields.io/badge/written%20in-go-29BEB0.svg)
![Latest GitHub release](https://img.shields.io/github/v/release/StiviiK/azctx.svg)
![GitHub stars](https://img.shields.io/github/stars/stiviik/azctx.svg?label=github%20stars)
![GitHub contributors](https://img.shields.io/github/contributors/stiviik/azctx.svg?label=github%20contributors)

**`azctx`** helps you switch between azure cli subscriptions back and forth:

![azctx demo GIF](img/azctx-demo.png)

# azctx(1)

```bash
azctx is a CLI tool for managing azure cli subscriptions.
        It is a helper for the azure cli and provides a simple interface for managing subscriptions.
        Pass a subscription name to select a specific subscription.
        Pass - to switch to the previous subscription.

Usage:
  azctx [- / NAME] [flags]
  azctx [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  login       Login to Azure
  version     Print the CLI version

Flags:
  -c, --current              Display the current active subscription
  -t, --fetch-tenant-names   Fetch the tenant names for the subscriptions
  -h, --help                 help for azctx
  -r, --refresh              Re-Authenticate and refresh the subscriptions

Use "azctx [command] --help" for more information about a command.
```

-----

## Installation

### Homebrew

* Install `azctx` with `brew install stiviik/tap/azctx`

### Linux

* Install the `azctx` binary from the [repository](https://github.com/StiviiK/azctx/releases)
* Add the `azctx` binary to your PATH

-----

## Troubleshooting

### Error: `AZURE_CONFIG_DIR is not set / a valid directory. [...]`

Run once `az configure` to create the configuration directory.

Check the [Microsoft Documentation](https://docs.microsoft.com/en-us/cli/azure/azure-cli-configuration?view=azure-cli-latest#cli-configuration-file) for the correct path to the azure cli config directory and export it as AZURE_CONFIG_DIR environment variable.  

```bash
export AZURE_CONFIG_DIR=$HOME/.azure
```

-----

## Todos

* [ ] Implement Unit Tests
