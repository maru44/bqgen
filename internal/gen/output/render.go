package output

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/maru44/bqgen/internal/core"
	"github.com/samber/lo"
	"golang.org/x/tools/imports"
)

//go:embed template.tmpl
var tmpl string

type dataWithPkg struct {
	Pkg  string
	Data *core.TableInput
}

func RenderTable(data *core.TableInput, pkg string) ([]byte, error) {
	funcMap := map[string]any{
		"gt":    core.GoTypeFromBQType,
		"pt":    core.TypeString,
		"pf":    parseField,
		"camel": camel,
	}
	t, err := template.New("bqgen").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	dp := dataWithPkg{
		Pkg:  pkg,
		Data: data,
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, dp); err != nil {
		return nil, fmt.Errorf("failed to render with template: %s", err.Error())
	}

	out, err := imports.Process("processing", buf.Bytes(), &imports.Options{
		FormatOnly: false,
		Comments:   true,
		TabIndent:  true,
		TabWidth:   8,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to process go-imports: %w", err)
	}

	return out, nil
}

func camel(s string) string {
	if len(s) == 0 {
		return s
	}
	return strcase.ToCamel(s)
}

func parseField(f *core.FieldInput) string {
	var elements []string
	elements = append(elements, fmt.Sprintf(`Name: "%s", Type: %s`, f.Name, core.TypeString(f.Type)))

	if f.Description != "" {
		elements = append(elements, fmt.Sprintf(`Description: "%s"`, f.Description))
	}
	if f.Required {
		elements = append(elements, "Required: true")
	}
	if f.Repeated {
		elements = append(elements, "Repeated: true")
	}
	if f.PolicyTags != nil {
		if f.PolicyTags == nil {
			elements = append(elements, "PolicyTags: &bigquery.PolicyTagList{}")
		}

		names := lo.Map(f.PolicyTags, func(v string, _ int) string {
			return fmt.Sprintf(`"%s"`, v)
		})
		elements = append(elements, fmt.Sprintf(`PolicyTags: &bigquery.PolicyTagList{Names: []string{%s}}`,
			strings.Join(names, ", ")))
	}
	if f.RecordedTableName != "" {
		elements = append(elements, fmt.Sprintf("Schema: %s.Schema", f.RecordedTableName))
	}
	if f.MaxLength != 0 {
		elements = append(elements, fmt.Sprintf("MaxLength: %d", f.MaxLength))
	}
	if f.Precision != 0 {
		elements = append(elements, fmt.Sprintf("Precision: %d", f.Precision))
	}
	if f.Scale != 0 {
		elements = append(elements, fmt.Sprintf("Scale: %d", f.Scale))
	}
	if f.DefaultValueExpression != "" {
		elements = append(elements, fmt.Sprintf(`DefaultValueExpression: "%s"`, f.DefaultValueExpression))
	}

	return strings.Join(elements, ", ")
}
