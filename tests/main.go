package tests

import "time"

//go:generate go run github.com/maru44/bqgen/generator

type SampleString string

type Person struct {
	Name  string `bigquery:"name"`
	Age   int    `bigquery:"age"`
	Sex   string `bigquery:"-"`
	Hobby string `bigquery:"hobby,nullable"`
}

type Animal struct {
	ID      string `bigquery:"id"`
	Kind    string `bigquery:"kind"`
	Goods   []*Good
	GoodPtr *Good `bigquery:"good_ptr"`
	GoodStr Good  `bigquery:"good_str"`
	strPtr  *string
	tim     time.Time
	timPtr  *time.Time
}

type Good struct {
	Name      string
	Sample    SampleString
	SamplePtr *SampleString
}
