package main

import (
	"cloud.google.com/go/bigquery"
)

var bqTypeWithoutZeroValueByUnderlyingType = map[string]bigquery.FieldType{
	"string":    bigquery.StringFieldType,
	"int":       bigquery.IntegerFieldType,
	"int8":      bigquery.IntegerFieldType,
	"int16":     bigquery.IntegerFieldType,
	"int32":     bigquery.IntegerFieldType,
	"int64":     bigquery.IntegerFieldType,
	"uint8":     bigquery.IntegerFieldType,
	"uint16":    bigquery.IntegerFieldType,
	"uint32":    bigquery.IntegerFieldType,
	"float32":   bigquery.FloatFieldType,
	"float64":   bigquery.FloatFieldType,
	"bool":      bigquery.BooleanFieldType,
	"[]byte":    bigquery.BytesFieldType,
	"time.Time": bigquery.TimeFieldType,
	// "civil.Date":     bigquery.DateFieldType,
	// "civil.Time":     bigquery.TimeFieldType,
	// "civil.DateTime": bigquery.DateTimeFieldType,
	// "*big.Rat":       bigquery.NumericFieldType,
}

var bqTypeByUnderlyingType = map[string]bigquery.FieldType{
	"string":    bigquery.StringFieldType,
	"int":       bigquery.IntegerFieldType,
	"int8":      bigquery.IntegerFieldType,
	"int16":     bigquery.IntegerFieldType,
	"int32":     bigquery.IntegerFieldType,
	"int64":     bigquery.IntegerFieldType,
	"uint8":     bigquery.IntegerFieldType,
	"uint16":    bigquery.IntegerFieldType,
	"uint32":    bigquery.IntegerFieldType,
	"float32":   bigquery.FloatFieldType,
	"float64":   bigquery.FloatFieldType,
	"bool":      bigquery.BooleanFieldType,
	"[]byte":    bigquery.BytesFieldType,
	"time.Time": bigquery.TimeFieldType,
	// "civil.Date":     bigquery.DateFieldType,
	// "civil.Time":     bigquery.TimeFieldType,
	// "civil.DateTime": bigquery.DateTimeFieldType,
	// "*big.Rat":       bigquery.NumericFieldType,
}
