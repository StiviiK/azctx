# `azctx`: Power tool for Azure CLI

#### Note: This repository is heavily orientated on [kubectx](https://github.com/ahmetb/kubectx). Most of the source code is from the original author of `kubectx`.

![Latest GitHub release](https://img.shields.io/github/release/stiviik/azctx.svg)
![GitHub stars](https://img.shields.io/github/stars/stiviik/azctx.svg?label=github%20stars)
![Proudly written in Bash](https://img.shields.io/badge/written%20in-bash-ff69b4.svg)

**`azctx`** helps you switch between subscriptions back and forth:
![azctx demo GIF](img/azctx-demo.gif)

# azctx(1)

azctx is a utility to switch between azure subscriptions.

```
USAGE:
  azctx                   : list the contexts
  azctx <NAME>            : switch to context <NAME>
  azctx -                 : switch to the previous context
  azctx -c, --current     : show the current context name
```

### Usage

```sh
$ azctx mvp
Switched to context "mvp"

$ kubectx -
Switched to context "workshop".

$ kubectx -
Switched to context "mvp".
```

`kubectx` supports <kbd>Tab</kbd> completion on bash/zsh/fish shells to help with
long context names. You don't have to remember full context names anymore.

-----

## Installation

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
AZURE_CONFIG_DIR=$HOME/.azure
```

-----

### Interactive mode

If you want `kubectx` and `kubens` commands to present you an interactive menu
with fuzzy searching, you just need to [install
`fzf`](https://github.com/junegunn/fzf) in your PATH.

![kubectx interactive search with fzf](img/kubectx-interactive.gif)

If you have `fzf` installed, but want to opt out of using this feature, set the environment variable `KUBECTX_IGNORE_FZF=1`.

---

#### Stargazers over time

[![Stargazers over time](https://starcharts.herokuapp.com/stiviik/azctx.svg)](https://starcharts.herokuapp.com/stiviik/azctx)
