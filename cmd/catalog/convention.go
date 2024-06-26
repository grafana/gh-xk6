package catalog

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v62/github"
	"github.com/grafana/gh-xk6/cmd/ghext"
	"github.com/spf13/cobra"
)

type conventionOptions struct {
	filename string
	force    bool
}

func conventionCmd() *cobra.Command {
	opts := new(conventionOptions)

	cmd := &cobra.Command{
		Use:   "convention",
		Short: "Create catalog based on convention",
		Long:  "Create k6 extension catalog from a convention based GitHub search",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return conventionRunE(cmd.Context(), opts)
		},
		Hidden: true,
	}

	flags := cmd.Flags()

	flags.BoolP("help", "h", false, "Help for "+cmd.Use+"command")
	flags.StringVarP(&opts.filename, "file", "f", defaultFilename, "Extension catalog filename")
	flags.BoolVar(&opts.force, "force", false, "Force overwriting of the existing file")

	return cmd
}

func conventionRunE(ctx context.Context, opts *conventionOptions) error {
	if err := checkExist(opts.filename, opts.force); err != nil {
		return err
	}

	mods, err := searchExtensions(ctx)
	if err != nil {
		return err
	}

	if err := updateVersions(ctx, mods); err != nil {
		return err
	}

	file, err := os.Create(filepath.Clean(opts.filename)) //nolint:forbidigo
	if err != nil {
		return err
	}

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	if err := enc.Encode(mods); err != nil {
		return err
	}

	return file.Close()
}

func searchExtensions(ctx context.Context) (catalogModules, error) {
	client := ghext.ContextGitHubClient(ctx)

	mods := make(catalogModules)

	page := 1

	for {
		result, resp, err := client.Search.Repositories(
			ctx,
			"topic:xk6",
			&github.SearchOptions{ListOptions: github.ListOptions{Page: page, PerPage: 100}},
		)
		if err != nil {
			return nil, err
		}

		for _, repo := range result.Repositories {
			if repo.GetArchived() || repo.GetIsTemplate() {
				continue
			}

			for _, name := range guessModuleName(repo) {
				mod := &catalogModule{Module: "github.com/" + repo.GetFullName()}

				// Do not overwrite an existing module.
				// The search returns results in descending order by star.
				// In this way, in the event of a module name conflict,
				// the one with more stars has higher priority.
				if _, has := mods[name]; !has {
					mods[name] = mod
				}
			}
		}

		if page >= resp.LastPage {
			break
		}

		page = resp.NextPage
	}

	return mods, nil
}

func guessModuleName(repo *github.Repository) []string {
	name := repo.GetName()

	var names []string

	add := func(value string) { names = append(names, value) }

	if !strings.HasPrefix(name, "xk6-") || hasTopic(repo, "xk6-related") || hasTopic(repo, "xk6-other") {
		return names
	}

	if topic, found := getTopic(repo, "xk6-output-"); found {
		add(strings.TrimPrefix(topic, "xk6-output-"))

		return names
	}

	if topic, found := getTopic(repo, "xk6-javascript-"); found {
		add(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.TrimPrefix(topic, "xk6-javascript-"),
					"-", "/",
				),
				"//", "-",
			),
		)

		return names
	}

	if strings.HasPrefix(name, "xk6-output-") {
		add(strings.TrimPrefix(name, "xk6-output-"))
	} else {
		add(name)
		add(strings.TrimPrefix(name, "xk6-"))
		add("k6/x/" + strings.TrimPrefix(name, "xk6-"))

		if idx := strings.LastIndex(name, "-"); idx >= 0 && idx < len(name) {
			add("k6/x/" + name[idx+1:])
		}
	}

	return names
}

func getTopic(repo *github.Repository, prefix string) (string, bool) {
	for _, topic := range repo.Topics {
		if strings.HasPrefix(topic, prefix) {
			return topic, true
		}
	}

	return "", false
}

func hasTopic(repo *github.Repository, value string) bool {
	for _, topic := range repo.Topics {
		if topic == value {
			return true
		}
	}

	return false
}
