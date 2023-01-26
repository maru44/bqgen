package main

import (
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

	for _, pk := range ps {
		p := newParser(pk)
		for _, f := range p.pkg.Syntax {
			structs := p.parseFile(f)
			pp.Println(structs)
		}

		// pp.Println(p.pkg.TypesInfo)
	}
}
