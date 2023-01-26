package main

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/k0kubun/pp"
	"golang.org/x/tools/go/packages"
)

type parser struct {
	pkg *packages.Package
}

// var primitiveType = map[string]struct{}{
// 	"string": {},
// 	"int": {},
// 	"int8": {},
// 	"int16": {},
// 	"int64": {},
// 	""
// }

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
		bqName, required, nullable, skip := p.getFileName(f)
		if skip {
			continue
		}

		ff := &field{
			name:     f.Names[0].Name,
			bqName:   bqName,
			required: required,
			nullable: nullable,
		}
		pp.Println(f.Type)
		if typ, ok := f.Type.(*ast.Ident); ok {
			ff.typ = typ.Name
			underlyingType := p.pkg.TypesInfo.TypeOf(typ).String()
			if ut, ok := bqTypeByUnderlyingType[underlyingType]; ok {
				ff.bqType = string(ut)
			}
		}
		sc.Fields = append(sc.Fields, ff)
	}
}

func (p *parser) isStringLikeType(ts *ast.TypeSpec) bool {
	t := p.pkg.TypesInfo.TypeOf(ts.Name)
	return t.Underlying().String() == "string"
}

func (p *parser) getFileName(f *ast.Field) (name string, required, nullable, skip bool) {
	required = true
	if f.Tag == nil {
		if len([]rune(f.Names[0].Name)) > 1 {
			name = strings.ToLower(string(f.Names[0].Name[0])) + f.Names[0].Name[1:]
		} else {
			name = strings.ToLower(f.Names[0].Name)
		}
		return
	}

	tagBq := tagReg.Find([]byte(f.Tag.Value))
	if len(tagBq) == 0 {
		if len([]rune(f.Names[0].Name)) > 1 {
			name = strings.ToLower(string(f.Names[0].Name[0])) + f.Names[0].Name[1:]
		} else {
			name = strings.ToLower(f.Names[0].Name)
		}
		return
	}

	tag := strings.TrimLeft(string(tagBq), `bigquery:`)
	tag = strings.Trim(tag, `"`)
	tagSlice := strings.Split(tag, ",")

	if len(tagSlice) > 2 {
		if tagSlice[1] == "nullable" {
			nullable = true
			required = false
		}
	}

	if tagSlice[0] == "-" {
		skip = true
		return
	}
	name = tagSlice[0]

	return
}
