package catalog

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v62/github"
	"github.com/grafana/gh-xk6/cmd/ghext"
	"github.com/jmespath/go-jmespath"

	"github.com/spf13/cobra"
)

var presets = map[string]string{ //nolint:gochecknoglobals
	"official":   "[?contains(tiers,'Official')]",
	"cloud":      "[?cloudEnabled == true]",
	"javascript": "[?contains(type,'JavaScript')]",
	"output":     "[?contains(type,'Output')]",
}

var errUnknownPreset = errors.New("unknown preset")

type preset string

func (p preset) String() string {
	return string(p)
}

func (p *preset) Set(value string) error {
	src, ok := presets[strings.ToLower(value)]
	if !ok {
		return errUnknownPreset
	}

	*p = preset(src)

	return nil
}

func (p preset) Type() string {
	return "name"
}

type filter string

func (f filter) String() string {
	return string(f)
}

func (f *filter) Set(value string) error {
	_, err := jmespath.Compile(value)
	if err != nil {
		return err
	}

	*f = filter(value)

	return nil
}

func (f filter) Type() string {
	return "query"
}

type importOptions struct {
	filter   filter
	preset   preset
	filename string
	force    bool
}

func (opts *importOptions) query() (*jmespath.JMESPath, error) {
	src := string(opts.filter)
	if len(src) == 0 {
		src = string(opts.preset)
	}

	if len(src) == 0 {
		src = "[*]"
	}

	return jmespath.Compile(src)
}

func importCmd() *cobra.Command {
	opts := new(importOptions)

	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import k6-docs extension registry",
		Long:  "Import k6-docs extension registry into k6 extension catalog",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return importRunE(cmd.Context(), opts)
		},
	}

	flags := cmd.Flags()

	flags.BoolP("help", "h", false, "Help for "+cmd.Use+"command")
	flags.VarP(&opts.filter, "filter", "q", "JMESPath query for filtering registry entries")
	flags.VarP(&opts.preset, "preset", "p", "Select a preset JMESPath filter by name")
	flags.StringVarP(&opts.filename, "file", "f", "k6catalog.json", "Extension catalog filename")
	flags.BoolVar(&opts.force, "force", false, "Force overwriting of the existing file")

	cmd.MarkFlagsMutuallyExclusive("filter", "preset")

	return cmd
}

func importRunE(ctx context.Context, opts *importOptions) error {
	if err := checkExist(opts.filename, opts.force); err != nil {
		return err
	}

	client := ghext.ContextGitHubClient(ctx)

	query, err := opts.query()
	if err != nil {
		return err
	}

	reg, err := downloadRegistry(ctx, client, query)
	if err != nil {
		return err
	}

	mods := reg.toModules()

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

func downloadRegistry(
	ctx context.Context, client *github.Client, query *jmespath.JMESPath,
) (*extensionRegistry, error) {
	content, _, _, err := client.Repositories.GetContents(
		ctx,
		"grafana",
		"k6-docs",
		"src/data/doc-extensions/extensions.json",
		nil,
	)
	if err != nil {
		return nil, err
	}

	str, err := content.GetContent()
	if err != nil {
		return nil, err
	}

	var bin []byte

	if query != nil {
		bin, err = applyFilter([]byte(str), query)
		if err != nil {
			return nil, err
		}
	} else {
		bin = []byte(str)
	}

	return parseExtensionRegistry(bin)
}
