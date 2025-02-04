package build

import "github.com/spf13/cobra"

// New returns a new catalog cobra command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build k6 with extensions",
		Long:  "Build k6 with one or more extensions",
		RunE:  runE,
	}

	return cmd
}

func runE(_ *cobra.Command, _ []string) error {
	return nil
}
