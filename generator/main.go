package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/k0kubun/pp"
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
	ps, _ := loadPackages()

	var schemas []*schema

	for _, pk := range ps {
		p := newParser(pk)
		for _, f := range p.pkg.Syntax {
			structs := p.parseFile(f)
			pp.Println(structs)
			schemas = append(schemas, structs...)
		}
	}

	data, err := render(schemas)
	if err != nil {
		log.Fatal("failed to render: ", err)
	}

	// TODO fix get by params
	modelPkgDir, err := filepath.Abs("./../out")
	if err != nil {
		log.Fatal("failed to get model package dir: ", err)
	}

	modelOutputPath := filepath.Join(modelPkgDir, "gen_modelgen.go")
	if err := os.WriteFile(modelOutputPath, data, 0600); err != nil {
		log.Fatal("failed to write file: ", err)
	}

}
