package core

import (
	"cloud.google.com/go/bigquery"
)

type (
	Schema struct {
		Schema    bigquery.Schema
		TypeByKey map[string]bigquery.FieldType
	}
)

func (s *Schema) Headers() []string {
	out := make([]string, len(s.Schema))
	for i, f := range s.Schema {
		out[i] = f.Name
	}
	return out
}

func (s *Schema) Type(fieldName string) (bigquery.FieldType, bool) {
	f, ok := s.TypeByKey[fieldName]
	return f, ok
}
