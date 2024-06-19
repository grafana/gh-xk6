package catalog

import (
	"encoding/json"

	"github.com/jmespath/go-jmespath"
)

type extensionRegistry struct {
	Extensions []registeredExtension `json:"extensions,omitempty"`
}

type registeredExtension struct {
	Name string   `json:"name,omitempty"`
	URL  string   `json:"url,omitempty"`
	Type []string `json:"type,omitempty"`
}

func applyFilter(src []byte, filter *jmespath.JMESPath) ([]byte, error) {
	loose := new(struct {
		Extensions interface{} `json:"extensions"`
	})

	if err := json.Unmarshal(src, loose); err != nil {
		return nil, err
	}

	data, err := filter.Search(loose.Extensions)
	if err != nil {
		return nil, err
	}

	loose.Extensions = data

	return json.Marshal(loose)
}

func parseExtensionRegistry(src []byte) (*extensionRegistry, error) {
	reg := new(extensionRegistry)

	if err := json.Unmarshal(src, reg); err != nil {
		return nil, err
	}

	return reg, nil
}
