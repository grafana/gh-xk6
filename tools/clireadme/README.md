# clireadme

A small library that helps to update the documentation in the README file of [cobra](https://github.com/spf13/cobra)-based CLI tools.

## Usage

Create a `tools/gendoc/main.go` file with the following content (the Main function must be called with your own cobra Command as a parameter).

Then run: `go run ./tools/gendoc README.md`

```go
// Package main contains CLI documentation generator tool.
package main

import (
	"github.com/grafana/xk6/tools/clireadme"
	"github.com/grafana/xk6/internal/sub/xk6"
)

func main() {
	clireadme.Main(xk6.New())
}

```