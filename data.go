package main

import "cloud.google.com/go/bigquery"

type (
	underlyingType string

	// field struct {
	// 	Name           string
	// 	BqName         string
	// 	Typ            string
	// 	UnderlyingType underlyingType
	// 	Array          bool
	// 	Ptr            bool
	// 	Required       bool
	// 	Nullable       bool
	// 	// Schema         *schema // if type is record
	// }

	dataset struct {
		ID        string
		ProjectID string
		Tables    []table
	}

	table struct {
		TableID   string
		DatasetID string
		ProjectID string
		Fields    []*field
	}

	field struct {
		Name                   string
		Description            string
		Repeated               bool
		Required               bool
		Type                   bigquery.FieldType
		PolicyTags             []string
		RecordedTableName      string
		MaxLength              int64
		Precision              int64
		Scale                  int64
		DefaultValueExpression string
	}
)

func (d *dataset) setDatasetIDToTable() {
	for _, t := range d.Tables {
		t.DatasetID = d.ID
		t.ProjectID = d.ProjectID
	}
}
