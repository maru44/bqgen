package input

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/maru44/bqgen/internal/core"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

var acceptableExt = map[string]struct{}{
	".yaml": {},
	".toml": {},
	".json": {},
}

func loadDefinitionDir(name string) ([][]core.DatasetInput, error) {
	dirs, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}
	var fileNames []string
	for _, d := range dirs {
		if _, ok := acceptableExt[filepath.Ext(d.Name())]; ok {
			fileNames = append(fileNames, d.Name())
		}
	}
	return loadDefinitionFiles(fileNames...)
}

func loadDefinitionFiles(names ...string) ([][]core.DatasetInput, error) {
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

func loadDefinitionFile(name string) ([]core.DatasetInput, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var out []core.DatasetInput
	ext := filepath.Ext(name)
	switch ext {
	case ".yaml":
		if err := yaml.Unmarshal(b, &out); err != nil {
			return nil, err
		}
	case ".toml":
		if err := toml.Unmarshal(b, &out); err != nil {
			return nil, err
		}
	case ".json":
		if err := json.Unmarshal(b, &out); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid file ext: %s", ext)
	}

	return out, nil
}
