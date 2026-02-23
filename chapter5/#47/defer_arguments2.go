package main

import "fmt"

const (
	StatusSuccess  = "success"
	StatusErrorFoo = "error_foo"
	StatusErrorBar = "error_bar"
)

type Struct struct {
	id string
}

func main() {
	s := &Struct{id: "foo"}
	defer s.print()
	s.id = "bar"
}

func (s *Struct) print() {
	fmt.Println(s.id)
}
