package catalog

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type createOptions struct {
	filename string
	force    bool
}

func createCmd() *cobra.Command {
	opts := new(createOptions)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create new extension catalog",
		Long:  "Create new extension catalog with only the mandatory k6 entry",
		RunE: func(c *cobra.Command, _ []string) error {
			return createRunE(c.Context(), opts)
		},
	}

	flags := cmd.Flags()

	flags.BoolP("help", "h", false, "Help for "+cmd.Use+"command")
	flags.StringVarP(&opts.filename, "file", "f", defaultFilename, "Extension catalog filename")
	flags.BoolVar(&opts.force, "force", false, "Force overwriting of the existing file")

	return cmd
}

func createRunE(ctx context.Context, opts *createOptions) error {
	if err := checkExist(opts.filename, opts.force); err != nil {
		return err
	}

	mods := catalogModules{"k6": &catalogModule{Module: k6ModulePath}}

	if err := updateVersions(ctx, mods); err != nil {
		return err
	}

	return mods.save(opts.filename)
}

func checkExist(filename string, force bool) error {
	if _, err := os.Stat(filename); (err == nil || !errors.Is(err, os.ErrNotExist)) && !force { //nolint:forbidigo
		return fmt.Errorf("%w: %s, use --force flag to overwrite it", errFileExists, filename)
	}

	return nil
}

var errFileExists = errors.New("file exists")
