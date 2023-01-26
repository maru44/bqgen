package main

type (
	underlyingType string

	schema struct {
		Name   string
		Fields []*field
	}

	field struct {
		Name           string
		BqName         string
		Typ            string
		UnderlyingType underlyingType
		Array          bool
		Ptr            bool
		Required       bool
		Nullable       bool
		// Schema         *schema // if type is record
	}
)
