# `azctx`: Power tool for the Azure CLI

![Latest GitHub release](https://img.shields.io/github/v/release/StiviiK/azctx.svg)
![GitHub stars](https://img.shields.io/github/stars/stiviik/azctx.svg?label=github%20stars)
![Proudly written in Go](https://img.shields.io/badge/written%20in-go-29BEB0.svg)

**`azctx`** helps you switch between subscriptions back and forth:
![azctx demo GIF](img/azctx-demo.gif)

# azctx(1)

azctx is a utility to switch between azure subscriptions.

```
USAGE:
  azctx                       : list the subscriptions
  azctx <NAME>                : switch to subscription <NAME>
  azctx -                     : switch to the previous subscription
  azctx -c, --current         : show the subscription name
  azctx -r, --refresh         : re-login and fetch all subscriptions
  azctx -h,--help             : show this message
```

### Usage

```sh
$ azctx mvp
Switched to context "mvp"

$ azctx -
Switched to context "workshop".

$ azctx -
Switched to context "mvp".
```

`azctx` supports <kbd>Tab</kbd> completion on bash/zsh/fish shells to help with
long context names. You don't have to remember full context names anymore. (WIP)

-----

## Installation

### Homebrew

```
brew install stiviik/tap/azctx
```

### Linux

TODO

-----
## Troubleshooting
### `AZURE_CONFIG_DIR` is not set
Find in the [Microsoft Documentation](https://docs.microsoft.com/en-us/cli/azure/azure-cli-configuration?view=azure-cli-latest#cli-configuration-file) the correct path to the azure cli config directory and set in manualy.  
For example set it in your `~/.bashrc`:
```
export AZURE_CONFIG_DIR=$HOME/.azure
```

Run once `az configure` to create the configuration directory.

-----
#### Stargazers over time

[![Stargazers over time](https://starcharts.herokuapp.com/stiviik/azctx.svg)](https://starcharts.herokuapp.com/stiviik/azctx)
