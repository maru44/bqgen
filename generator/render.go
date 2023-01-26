package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	"cloud.google.com/go/bigquery"
	"github.com/iancoleman/strcase"
	"golang.org/x/tools/imports"
)

//go:embed template.tmpl
var tmpl string

func render(data []*schema) ([]byte, error) {
	funcMap := map[string]any{
		"bqType":         bqType,
		"camel":          camel,
		"schemaRelation": schemaRelation,
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

func bqType(underlyingType string) string {
	if bt, ok := bqTypeByUnderlyingType[underlyingType]; ok {
		return parseType(bt)
	}
	return parseType(bigquery.RecordFieldType)
}

func parseType(s bigquery.FieldType) string {
	switch s {
	case bigquery.StringFieldType:
		return "bigquery.StringFieldType"
	case bigquery.BytesFieldType:
		return "bigquery.StringFieldType"
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
	}
	panic(fmt.Sprintf("no such type: %s", s))
}

func camel(s string) string {
	if len(s) == 0 {
		return s
	}
	return strcase.ToCamel(s)
}

func schemaRelation(f *field) string {
	if _, ok := bqTypeByUnderlyingType[f.UnderlyingType]; ok {
		return ""
	}
	return fmt.Sprintf(", Schema: %s.Schema", camel(f.Typ))
}