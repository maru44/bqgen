package main

type structDefine struct {
	Name   string
	Fields []*field
}

type field struct {
	name     string
	bqName   string
	typ      string
	required bool
	nullable bool
}
