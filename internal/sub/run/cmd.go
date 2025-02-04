package run

import "github.com/spf13/cobra"

// New returns a new catalog cobra command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run k6 with extensions",
		Long:  "Run k6 with actual extension",
		RunE:  runE,
	}

	return cmd
}

func runE(_ *cobra.Command, _ []string) error {
	return nil
}
