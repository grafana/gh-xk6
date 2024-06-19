// Package cmd contains cobra commands of gh xk6.
package cmd

import (
	"context"

	"github.com/grafana/gh-xk6/cmd/catalog"
	"github.com/grafana/gh-xk6/cmd/ghext"
	"github.com/spf13/cobra"
)

var version = "dev"

// New creates new cobra command for gh-xk6 command.
func New() *cobra.Command {
	root := &cobra.Command{
		Use:   "xk6",
		Short: "Maintain k6 extensions hosted on GitHub",
		Long: `A gh extension that helps maintain k6 extensions hosted on GitHub.
`,
		SilenceUsage:      true,
		SilenceErrors:     true,
		Version:           version,
		Annotations:       map[string]string{cobra.CommandDisplayNameAnnotation: "gh xk6"},
		CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	}

	root.SetContext(ghext.Context(context.TODO()))

	root.SetVersionTemplate(`gh-xk6 {{printf "version %s" .Version}}
`)

	flags := root.Flags()

	flags.BoolP("version", "V", false, "version for gh xk6")

	root.AddCommand(catalog.New())

	return root
}
