// Package main contains CLI documentation generator tool.
package main

import (
	"github.com/grafana/xk6/internal/sub/xk6"
	"github.com/grafana/xk6/tools/clireadme"
)

func main() {
	clireadme.Main(xk6.New(), 1)
}
