package catalog

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v62/github"
	"github.com/grafana/gh-xk6/cmd/ghext"
	"github.com/spf13/cobra"
)

type updateOptions struct {
	filename string
}

func updateCmd() *cobra.Command {
	opts := new(updateOptions)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update versions in extension catalog",
		Long:  "Update versions in extension catalog using GitHub search",
		RunE: func(c *cobra.Command, _ []string) error {
			return updateRunE(c.Context(), opts)
		},
	}

	flags := cmd.Flags()

	flags.BoolP("help", "h", false, "Help for "+cmd.Use+"command")
	flags.StringVarP(&opts.filename, "file", "f", defaultFilename, "Extension catalog filename")

	return cmd
}

func updateRunE(ctx context.Context, opts *updateOptions) error {
	bin, err := os.ReadFile(opts.filename) //nolint:forbidigo
	if err != nil {
		return err
	}

	var mods catalogModules

	if err := json.Unmarshal(bin, &mods); err != nil {
		return err
	}

	if err := updateVersions(ctx, mods); err != nil {
		return err
	}

	return mods.save(opts.filename)
}

func updateVersions(ctx context.Context, mods catalogModules) error {
	client := ghext.ContextGitHubClient(ctx)

	for name, mod := range mods {
		var owner, repo string

		if name == "k6" {
			owner = "grafana"
			repo = "k6"
		} else {
			parts := strings.SplitN(mod.Module, "/", 4)

			owner = parts[1]
			repo = parts[2]
		}

		mod.Versions = nil

		tags, _, err := client.Repositories.ListTags(
			ctx,
			owner,
			repo,
			&github.ListOptions{PerPage: 100},
		)
		if err != nil {
			return err
		}

		if len(tags) == 0 {
			continue
		}

		for _, tag := range tags {
			name := tag.GetName()
			if name[0] != 'v' {
				continue
			}

			ver, err := semver.NewVersion(name)
			if err != nil {
				continue
			}

			mod.Versions = append(mod.Versions, ver)
		}
	}

	return nil
}
