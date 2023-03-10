// Code generated by "github.com/maru44/bqgen"; DO NOT EDIT.

package out

import (
	"github.com/maru44/bqgen/model"

	"cloud.google.com/go/bigquery"
)

type SampleTable01 struct {
	Id   string `json:"id" toml:"id" yaml:"id" bigquery:"id"`
	Name string `json:"name" toml:"name" yaml:"name" bigquery:"name"`
}

var SampleTable01Schema = &model.Table{
	ID:        "sample_table_01",
	DatasetID: "test_dataset",
	Schema: bigquery.Schema{
		{Name: "id", Type: bigquery.StringFieldType, Required: true},
		{Name: "name", Type: bigquery.StringFieldType, Description: "this is name", Scale: 4},
	},
	TypeByFieldName: map[string]bigquery.FieldType{
		"id":   bigquery.StringFieldType,
		"name": bigquery.StringFieldType,
	},
}

var SampleTable01Columns = struct {
	Id   string
	Name string
}{
	Id:   "id",
	Name: "name",
}
