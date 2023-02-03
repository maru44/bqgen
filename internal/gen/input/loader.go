package input

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/maru44/bqgen/internal/core"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

// TODO impl
func LoadDefinition() {}

// TODO load dir

// TODO private
func LoadDefinitionFiles(names ...string) ([][]core.DatasetInput, error) {
	out := make([][]core.DatasetInput, len(names))
	for i, n := range names {
		var err error
		out[i], err = loadDefinitionFile(n)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// TODO other mimetype
func loadDefinitionFile(name string) ([]core.DatasetInput, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var out []core.DatasetInput
	switch {
	case strings.HasSuffix(name, ".yaml"):
		if err := yaml.Unmarshal(b, &out); err != nil {
			return nil, err
		}
	case strings.HasSuffix(name, ".toml"):
		if err := toml.Unmarshal(b, &out); err != nil {
			return nil, err
		}
	case strings.HasSuffix(name, ".json"):
		if err := json.Unmarshal(b, &out); err != nil {
			return nil, err
		}
	default:
		// TODO
	}

	return out, nil
}
