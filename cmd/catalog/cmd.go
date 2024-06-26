package catalog

import (
	"github.com/spf13/cobra"
)

// New returns a new catalog cobra command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "catalog",
		Short: "Maintain k6 extension catalog",
		Long:  "Maintain k6 extension catalog based on GitHub search",
	}

	cmd.AddCommand(createCmd(), importCmd(), updateCmd(), conventionCmd())

	return cmd
}
