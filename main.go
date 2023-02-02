package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var tagReg *regexp.Regexp

func init() {
	var err error
	tagReg, err = regexp.Compile(`bigquery:"[a-zA-Z0-9-_,]+"`)
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := loadBqDataFromConfig()
	if err != nil {
		log.Fatal("failed to load: ", err)
	}

	// TODO fix get by params
	modelPkgDir, err := filepath.Abs("./../out")
	if err != nil {
		log.Fatal("failed to get model package dir: ", err)
	}

	for _, d := range data {
		d.setDatasetIDToTable()
		for _, t := range d.Tables {
			modelDdata, err := render(t)
			if err != nil {
				log.Fatal("failed to render model data: ", err)
			}

			modelOutputPath := filepath.Join(modelPkgDir, fmt.Sprintf("%s.go", t.TableID))
			if err := os.WriteFile(modelOutputPath, modelDdata, 0600); err != nil {
				log.Fatal("failed to write file: ", err)
			}
		}
	}

}
