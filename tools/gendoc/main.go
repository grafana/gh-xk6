// Package main contains CLI documentation generator tool.
package main

import (
	_ "embed"

	"github.com/grafana/clireadme"
	"github.com/grafana/gh-xk6/cmd"
)

func main() {
	root := cmd.New()
	clireadme.Main(root, 1)
}
