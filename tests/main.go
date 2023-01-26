package tests

//go:generate go run github.com/maru44/bqlgen/generator

type Person struct {
	Name  string `bigquery:"name"`
	Age   int    `bigquery:"age"`
	Sex   string
	Hobby string `bigquery:"hobby,nullable"`
}

type Animal struct {
	ID   string `bigquery:"id"`
	Kind string `bigquery:"kind"`
}
