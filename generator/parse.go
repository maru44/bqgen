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

func (p *parser) parseFile(file *ast.File) []*schema {
	var schemas []*schema
	for _, decl := range file.Decls {
		switch it := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range it.Specs {
				// pp.Println(spec)

				switch ts := spec.(type) {
				case *ast.TypeSpec:
					sc := &schema{
						Name: ts.Name.Name,
					}

					if st, ok := ts.Type.(*ast.StructType); ok {
						p.parseStruct(st, sc)
					}
					schemas = append(schemas, sc)
				}
			}
		}
	}

	return schemas
}

func (p *parser) parseStruct(st *ast.StructType, sc *schema) {
	for _, f := range st.Fields.List {
		if f.Tag == nil {
			continue
		}

		tagBq := tagReg.Find([]byte(f.Tag.Value))
		if len(tagBq) == 0 {
			continue
		}
		tag := strings.TrimLeft(string(tagBq), `bigquery:`)
		tag = strings.Trim(tag, `"`)
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
		sc.Fields = append(sc.Fields, ff)
	}
}

func (p *parser) isStringLikeType(ts *ast.TypeSpec) bool {
	t := p.pkg.TypesInfo.TypeOf(ts.Name)
	return t.Underlying().String() == "string"
}
