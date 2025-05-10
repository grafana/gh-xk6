# gh-xk6

**GitHub CLI extension for xk6 subcommand**

The **gh-xk6** extension allows [grafana/xk6](https://github.com/grafana/xk6) commands to be used from the [GitHub CLI](https://cli.github.com/).

## Install

The extension can be installed using the command:

```bash
gh extension install grafana/gh-xk6
```

## Usage

The extension can be used in the same way as the `xk6` command, with the difference that the command line must be started with the `gh` command.

The following command will build k6 with the [xk6-example](https://github.com/grafana/xk6-example) extension:

```bash
gh xk6 build --with github.com/grafana/xk6-example
```
