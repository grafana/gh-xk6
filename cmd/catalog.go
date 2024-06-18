package cmd

import "github.com/spf13/cobra"

func catalogCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "catalog",
		Short: "Generate k6 extension catalog",
		Long:  "Generate k6 extension catalog based on GitHub topic search",
		RunE: func(_ *cobra.Command, _ []string) error {
			return nil
		},
	}

	return cmd
}
