package tests

//go:generate go run github.com/maru44/bqgen/generator

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
	Good    *Good
	GoodStr Good
	strPtr  *string
}

type Good struct {
	Name string
}
