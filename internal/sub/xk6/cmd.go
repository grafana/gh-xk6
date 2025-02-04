// Package xk6 contains cobra commands of xk6.
package xk6

import (
	"context"

	"github.com/grafana/xk6/internal/sub/build"
	"github.com/grafana/xk6/internal/sub/run"
	"github.com/spf13/cobra"
)

// New creates new cobra command for gh-xk6 command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "xk6",
		Short: "Maintain k6 extensions hosted on GitHub",
		Long: `A gh extension that helps maintain k6 extensions hosted on GitHub.
`,
		SilenceUsage:      true,
		SilenceErrors:     true,
		CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	}

	cmd.SetContext(Context(context.TODO()))

	cmd.AddCommand(build.New(), run.New())

	return cmd
}
