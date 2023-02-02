package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"cloud.google.com/go/bigquery"
	"github.com/iancoleman/strcase"
	"github.com/samber/lo"
	"golang.org/x/tools/imports"
)

//go:embed template.tmpl
var tmpl string

func render(data table) ([]byte, error) {
	funcMap := map[string]any{
		// "bqType": bqTypeStr,
		"pt":    parseType,
		"pf":    parseField,
		"camel": camel,
	}
	t, err := template.New("bqgen").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
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

// func bqTypeStr(ut underlyingType) string {
// 	return parseType(bqType(ut))
// }

// func bqType(ut underlyingType) bigquery.FieldType {
// 	if bt, ok := bqTypeByUnderlyingType[ut]; ok {
// 		return bt
// 	}
// 	return bigquery.RecordFieldType
// }

func camel(s string) string {
	if len(s) == 0 {
		return s
	}
	return strcase.ToCamel(s)
}

func parseField(f *field) string {
	var elements []string
	elements = append(elements, fmt.Sprintf(`Name: "%s", Type: %s`, f.Name, parseType(f.Type)))

	if f.Description != "" {
		elements = append(elements, fmt.Sprintf("Description: %s", f.Description))
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
		elements = append(elements, fmt.Sprintf("DefaultValueExpression: %s", f.DefaultValueExpression))
	}

	return strings.Join(elements, ", ")
}

// TODO ignore case
func parseType(s bigquery.FieldType) string {
	switch s {
	case bigquery.StringFieldType:
		return "bigquery.StringFieldType"
	case bigquery.BytesFieldType:
		return "bigquery.BytesFieldType"
	case bigquery.IntegerFieldType, "INT64":
		return "bigquery.IntegerFieldType"
	case bigquery.FloatFieldType, "FLOAT64":
		return "bigquery.FloatFieldType"
	case bigquery.TimestampFieldType:
		return "bigquery.TimestampFieldType"
	case bigquery.RecordFieldType, "STRUCT":
		return "bigquery.RecordFieldType"
	case bigquery.DateFieldType:
		return "bigquery.DateFieldType"
	case bigquery.TimeFieldType:
		return "bigquery.TimeFieldType"
	case bigquery.DateTimeFieldType:
		return "bigquery.DateTimeFieldType"
	case bigquery.NumericFieldType, "DECIMAL":
		return "bigquery.NumericFieldType"
	case bigquery.BigNumericFieldType, "BIGDECIMAL":
		return "bigquery.BigNumericFieldType"
	case bigquery.GeographyFieldType:
		return "bigquery.GeographyFieldType"
	case bigquery.IntervalFieldType:
		return "bigquery.IntervalFieldType"
	case bigquery.BooleanFieldType, "BOOL":
		return "bigquery.BooleanFieldType"
	case bigquery.JSONFieldType:
		return "bigquery.JSONFieldType"
	default:
		return string(s)
	}
	// panic(fmt.Sprintf("no such type: %s", s))
}

// func bqOriginType(in string) bool {
// 	switch in {
// 	case "bigquery.StringFieldType":
// 	case "bigquery.BytesFieldType":
// 	case "bigquery.IntegerFieldType":
// 	case "bigquery.FloatFieldType":
// 	case "bigquery.TimestampFieldType":
// 	case "bigquery.RecordFieldType":
// 	case "bigquery.DateFieldType":
// 	case "bigquery.TimeFieldType":
// 	case "bigquery.DateTimeFieldType":
// 	case "bigquery.NumericFieldType":
// 	case "bigquery.BigNumericFieldType":
// 	case "bigquery.GeographyFieldType":
// 	case "bigquery.IntervalFieldType":
// 	case "bigquery.BooleanFieldType":
// 	case "bigquery.JSONFieldType":
// 		return true
// 	}
// 	return false
// }
