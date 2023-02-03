package core

import "cloud.google.com/go/bigquery"

type (
	DatasetInput struct {
		ID        string
		ProjectID string
		Tables    []*TableInput
	}

	TableInput struct {
		ID        string
		DatasetID string
		ProjectID string
		Fields    []*FieldInput
	}

	FieldInput struct {
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

func (d *DatasetInput) SetDatasetIDToTable() {
	for _, t := range d.Tables {
		t.DatasetID = d.ID
		t.ProjectID = d.ProjectID
	}
}
