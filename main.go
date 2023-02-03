package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/maru44/bqgen/internal/gen/input"
	"github.com/maru44/bqgen/internal/gen/output"
)

func main() {
	ctx := context.Background()
	cfg, err := input.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}
	if err := cfg.Validate(ctx); err != nil {
		log.Fatal("invalid config:", err)
	}

	// TODO fix func
	data, err := input.LoadDefinitionFiles(cfg.Files...)
	if err != nil {
		log.Fatal("failed to load: ", err)
	}

	// TODO fix get by params
	modelPkgDir, err := filepath.Abs(cfg.Output)
	if err != nil {
		log.Fatal("failed to get model package dir: ", err)
	}

	for _, d := range data {
		for _, dd := range d {
			dd.SetDatasetIDToTable()
			for _, t := range dd.Tables {
				modelDdata, err := output.RenderTable(t, cfg.Pkg)
				if err != nil {
					log.Fatal("failed to render model data: ", err)
				}

				modelOutputPath := filepath.Join(modelPkgDir, t.ID+".go")
				if err := os.WriteFile(modelOutputPath, modelDdata, 0600); err != nil {
					log.Fatal("failed to write file: ", err)
				}
			}
		}
	}

}
