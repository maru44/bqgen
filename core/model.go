package core

import (
	"cloud.google.com/go/bigquery"
)

type (
	Table struct {
		// ProjectID       string
		ID              string
		DatasetID       string
		Schema          bigquery.Schema
		TypeByFieldName map[string]bigquery.FieldType
	}
)

func (s *Table) Fields() []string {
	out := make([]string, len(s.Schema))
	for i, f := range s.Schema {
		out[i] = f.Name
	}
	return out
}

func (s *Table) Type(fieldName string) (bigquery.FieldType, bool) {
	f, ok := s.TypeByFieldName[fieldName]
	return f, ok
}
