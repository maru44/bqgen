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

					// pp.Println(ts.Type)

					if st, ok := ts.Type.(*ast.StructType); ok {
						// if len(st.Fields.List) == 0 {
						// 	continue
						// }
						p.parseStruct(st, sc)
						schemas = append(schemas, sc)
					}
				}
			}
		}
	}

	return schemas
}

func (p *parser) parseStruct(st *ast.StructType, sc *schema) {
	for _, f := range st.Fields.List {
		ff, ok := p.parseFields(f)
		if !ok {
			continue
		}

		sc.Fields = append(sc.Fields, ff)
	}
}

func (p *parser) isStringLikeType(ts *ast.TypeSpec) bool {
	t := p.pkg.TypesInfo.TypeOf(ts.Name)
	return t.Underlying().String() == "string"
}

func (p *parser) parseFields(f *ast.Field) (*field, bool) {
	// pp.Println(f) // TODO del debug
	bqName, nullable, skip := p.getFileName(f)
	if skip {
		return nil, false
	}

	out := &field{
		Name:     f.Names[0].Name,
		BqName:   bqName,
		Nullable: nullable,
	}

	switch typ := f.Type.(type) {
	case *ast.Ident:
		out.Typ, out.UnderlyingType = p.parseIdent(typ)
	case *ast.ArrayType:
		out.Array = true
		switch elt := typ.Elt.(type) {
		case *ast.StarExpr:
			out.Ptr = true
			switch x := elt.X.(type) {
			case *ast.Ident:
				out.Typ, out.UnderlyingType = p.parseIdent(x)
			case *ast.SelectorExpr:
				out.Typ = x.Sel.Name
				// out.underlyingType = strings.TrimLeft(p.pkg.TypesInfo.TypeOf(typ).String(), "*")
				out.UnderlyingType = p.pkg.TypesInfo.TypeOf(typ).String()
			}
		case *ast.Ident:
			out.Typ, out.UnderlyingType = p.parseIdent(elt)
		}
	case *ast.StarExpr:
		out.Ptr = true
		if x, ok := typ.X.(*ast.Ident); ok {
			out.Typ, out.UnderlyingType = p.parseIdent(x)
		}
		switch x := typ.X.(type) {
		case *ast.Ident:
			out.Typ, out.UnderlyingType = p.parseIdent(x)
		case *ast.SelectorExpr:
			out.Typ = x.Sel.Name
			out.UnderlyingType = p.pkg.TypesInfo.TypeOf(typ).String()
		}
	case *ast.SelectorExpr:
		// if x, ok := typ.X.(*ast.Ident); ok {
		// 	if x.Obj == nil {
		// 		out.typ = x.Name
		// 	}
		// }
		// out.typ += "." + typ.Sel.Name
		out.Typ = typ.Sel.Name
		out.UnderlyingType = p.pkg.TypesInfo.TypeOf(typ).String()
	}
	if _, ok := bqTypeWithoutZeroValueByUnderlyingType[out.UnderlyingType]; ok && !out.Nullable {
		out.Required = true
	}
	return out, true
}

func (p *parser) parseIdent(ide *ast.Ident) (string, string) {
	if ide.Obj == nil {
		return ide.Name, p.pkg.TypesInfo.TypeOf(ide).String()
	}
	if ide.Obj.Decl != nil {
		if spec, ok := ide.Obj.Decl.(*ast.TypeSpec); ok {
			switch typ := spec.Type.(type) {
			case *ast.Ident:
				// enum like
				return ide.Name, p.pkg.TypesInfo.TypeOf(typ).String()
			case *ast.StructType:
				// pass
			}
		}
	}
	return ide.Obj.Name, p.pkg.TypesInfo.TypeOf(ide).String()
}

func (p *parser) getFileName(f *ast.Field) (name string, nullable, skip bool) {
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
		}
	}

	if tagSlice[0] == "-" {
		skip = true
		return
	}
	name = tagSlice[0]

	return
}
