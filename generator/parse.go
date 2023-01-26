package main

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/packages"
)

type parser struct {
	pkg *packages.Package
}

func newParser(pkg *packages.Package) *parser {
	return &parser{
		pkg: pkg,
	}
}

func loadPackages() ([]*packages.Package, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to load package: %w", err)
	}
	return pkgs, err
}

func (p *parser) parseFile(file *ast.File) []*structDefine {
	var structs []*structDefine
	for _, decl := range file.Decls {
		switch it := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range it.Specs {
				// pp.Println(spec)

				switch ts := spec.(type) {
				case *ast.TypeSpec:
					sd := &structDefine{
						Name: ts.Name.Name,
					}

					if st, ok := ts.Type.(*ast.StructType); ok {
						p.parseStruct(st, sd)
					}
					structs = append(structs, sd)
				}
			}
		}
	}

	return structs
}

func (p *parser) parseStruct(st *ast.StructType, def *structDefine) {
	var isBq bool
	for _, f := range st.Fields.List {
		if f.Tag == nil {
			continue
		}
		tags := strings.Split(f.Tag.Value, " ")

		for _, t := range tags {
			if strings.HasPrefix(t, "`bigquery") {
				tagB := tagReg.Find([]byte(t))
				tag := strings.Trim(string(tagB), `"`)
				fmt.Println(tag)

				tagSlice := strings.Split(tag, ",")

				var required bool = true
				var nullble bool
				if len(tagSlice) > 2 {
					if tagSlice[1] == "nullable" {
						nullble = true
						required = false
					}
				}

				ff := &field{
					name:     f.Names[0].Name,
					bqName:   tagSlice[0],
					required: required,
					nullable: nullble,
				}
				if typ, ok := f.Type.(*ast.Ident); ok {
					ff.typ = typ.Name
				}
				def.Fields = append(def.Fields, ff)
				isBq = true
				break
			}
		}
		if !isBq {
			continue
		}
	}
}

func (p *parser) isStringLikeType(ts *ast.TypeSpec) bool {
	t := p.pkg.TypesInfo.TypeOf(ts.Name)
	return t.Underlying().String() == "string"
}
