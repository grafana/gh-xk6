package catalog

import (
	"bytes"
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Masterminds/semver/v3"
)

type catalogModule struct {
	Module   string            `json:"module,omitempty"`
	Versions []*semver.Version `json:"versions,omitempty"`
}

func (mod catalogModule) MarshalJSON() ([]byte, error) {
	data := struct {
		Module   string   `json:"module,omitempty"`
		Versions []string `json:"versions,omitempty"`
	}{Module: mod.Module}

	var vers []*semver.Version

	vers = append(vers, mod.Versions...)
	sort.Slice(data.Versions, func(i, j int) bool {
		return vers[i].GreaterThan(vers[j])
	})

	for _, ver := range vers {
		data.Versions = append(data.Versions, "v"+ver.String())
	}

	return json.Marshal(data)
}

type catalogModules map[string]*catalogModule

func (mods catalogModules) MarshalJSON() ([]byte, error) {
	keys := make([]string, 0, len(mods))

	for key := range mods {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var buff bytes.Buffer

	buff.WriteRune('{')

	enc := json.NewEncoder(&buff)

	for idx, key := range keys {
		if idx > 0 {
			buff.WriteRune(',')
		}

		val := mods[key]

		if err := enc.Encode(key); err != nil {
			return nil, err
		}

		buff.WriteRune(':')

		if err := enc.Encode(val); err != nil {
			return nil, err
		}
	}

	buff.WriteRune('}')

	return buff.Bytes(), nil
}

func (mods catalogModules) save(filename string) error {
	file, err := os.Create(filepath.Clean(filename)) //nolint:forbidigo
	if err != nil {
		return err
	}

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	if err = enc.Encode(mods); err != nil {
		return err
	}

	return file.Close()
}

func (reg *extensionRegistry) toModules() catalogModules {
	mods := make(catalogModules, len(reg.Extensions)+1)

	add := func(path, name string) {
		mod := &catalogModule{
			Module: path,
		}
		mods[name] = mod
	}

	add(k6ModulePath, "k6")

	for _, regExt := range reg.Extensions {
		loc, err := url.Parse(regExt.URL)
		if err != nil {
			continue
		}

		path := loc.Host + loc.Path

		for _, typ := range regExt.Type {
			if typ == "Output" {
				add(path, regExt.Name)
				add(path, strings.TrimPrefix(regExt.Name, "xk6-output-"))
				add(path, strings.TrimPrefix(regExt.Name, "xk6-"))
			}

			if typ == "JavaScript" {
				add(path, "k6/x/"+strings.TrimPrefix(regExt.Name, "xk6-"))

				if idx := strings.LastIndex(regExt.Name, "-"); idx >= 0 && idx < len(regExt.Name) {
					add(path, "k6/x/"+regExt.Name[idx+1:])
				}
			}
		}
	}

	return mods
}

const (
	k6ModulePath    = "go.k6.io/k6"
	defaultFilename = "k6catalog.json"
)
