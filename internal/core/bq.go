package core

import (
	"fmt"

	"cloud.google.com/go/bigquery"
)

// TODO ignore case
func TypeString(s bigquery.FieldType) string {
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
	}
	panic(fmt.Sprintf("no such type: %s", s))
}

func GoTypeFromBQType(s bigquery.FieldType) string {
	switch s {
	case bigquery.StringFieldType:
		return "string"
	case bigquery.BytesFieldType:
		return "[]byte"
	case bigquery.IntegerFieldType, "INT64":
		return "int64"
	case bigquery.FloatFieldType, "FLOAT64":
		return "float64"
	case bigquery.TimestampFieldType:
		return "time.Time"
	case bigquery.RecordFieldType, "STRUCT":
		panic(fmt.Sprintf("must not reach here: %s", s))
	case bigquery.DateFieldType:
		return "civil.Date"
	case bigquery.TimeFieldType:
		return "civil.Time"
	case bigquery.DateTimeFieldType:
		return "civil.DateTime"
	case bigquery.NumericFieldType, "DECIMAL",
		bigquery.BigNumericFieldType, "BIGDECIMAL":
		return "*big.Rat"
	case bigquery.GeographyFieldType:
		return "string"
	case bigquery.IntervalFieldType:
		return "time.Duration"
	case bigquery.BooleanFieldType, "BOOL":
		return "bool"
	case bigquery.JSONFieldType:
		return "types.JSON"
	}
	panic(fmt.Sprintf("no such type: %s", s))
}
