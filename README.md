# `azctx`: Power tool for the Azure CLI

#### Note: This repository is heavily orientated on [kubectx](https://github.com/ahmetb/kubectx). Most of the source code is from the original author of `kubectx`.

![Latest GitHub release](https://img.shields.io/github/v/release/StiviiK/azctx.svg)
![GitHub stars](https://img.shields.io/github/stars/stiviik/azctx.svg?label=github%20stars)
![Proudly written in Bash](https://img.shields.io/badge/written%20in-bash-ff69b4.svg)

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
brew tap stiviik/tap
brew install azctx
```

### Linux

Since `azctx` is written in Bash, you should be able to install
them to any POSIX environment that has Bash installed.

- Download the `azctx` script.
- Either:
  - save them all to somewhere in your `PATH`,
  - or save them to a directory, then create a symlink to `azctx` from
    somewhere in your `PATH`, like `/usr/local/bin`
- Make `azctx` executable (`chmod +x ...`)

Example installation steps:

``` bash
sudo git clone https://github.com/StiviiK/azctx /opt/azctx
sudo ln -s /opt/azctx/azctx /usr/local/bin/azctx
```

## Troubleshooting
### `AZURE_CONFIG_DIR` is not set
Find in the [Microsoft Documentation](https://docs.microsoft.com/en-us/cli/azure/azure-cli-configuration?view=azure-cli-latest#cli-configuration-file) the correct path to the azure cli config directory and set in manualy.  
For example set it in your `~/.bashrc`:
```
export AZURE_CONFIG_DIR=$HOME/.azure
```

Run once `az configure` to create the configuration directory.

-----

### Interactive mode

If you want `azctx` command to present you an interactive menu
with fuzzy searching, you just need to [install
`fzf`](https://github.com/junegunn/fzf) in your PATH.

![azctx interactive search with fzf](img/azctx-interactive.gif)

If you have `fzf` installed, but want to opt out of using this feature, set the environment variable `AZCTX_IGNORE_FZF=1`.

---

#### Stargazers over time

[![Stargazers over time](https://starcharts.herokuapp.com/stiviik/azctx.svg)](https://starcharts.herokuapp.com/stiviik/azctx)
