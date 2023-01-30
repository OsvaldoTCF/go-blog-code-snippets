package main

import "fmt"

type FirstTypes interface {
	~int | ~string
}

type First[T FirstTypes] struct {
	v T
}

func (f First[T]) Value() T {
	return f.v
}

func main() {
	fmt.Println(First[int]{v: 1}.Value())
	fmt.Println(First[string]{v: "a"}.Value())
	// wrong type
	fmt.Println(First[bool]{v: true}.Value())
}
