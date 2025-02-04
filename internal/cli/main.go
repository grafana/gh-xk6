// Package cli contains common Main function for xk6 CLI incarnations.
package cli

import (
	"log"

	"github.com/grafana/xk6/internal/sub/xk6"
	"github.com/spf13/cobra"
)

// Main is a common main function for various xk6 executables (xk6, gh-xk6).
func Main(appname string) {
	root := xk6.New()

	root.Version = version

	if len(appname) > 0 {
		root.Annotations = map[string]string{cobra.CommandDisplayNameAnnotation: appname}
	}

	if err := root.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
