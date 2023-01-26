package main

type (
	schema struct {
		Name   string
		Fields []*field
	}

	field struct {
		Name           string
		BqName         string
		Typ            string
		UnderlyingType string
		Array          bool
		Ptr            bool
		Required       bool
		Nullable       bool
		// Schema         *schema // if type is record
	}
)
