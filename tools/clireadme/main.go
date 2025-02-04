package clireadme

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// Main updates the markdown documentation recursively based on cobra Command.
func Main(root *cobra.Command, headingOffset int) {
	const nArgs = 2

	exe := filepath.Base(os.Args[0])
	if len(os.Args) != nArgs { //nolint:gomnd
		fmt.Fprintf(os.Stderr, "usage: %s filename", exe)
		os.Exit(1)
	}

	if err := Update(root, os.Args[1], headingOffset); err != nil {
		fmt.Fprintf(os.Stderr, "%s: error: %s\n", exe, err)
		os.Exit(1)
	}
}
