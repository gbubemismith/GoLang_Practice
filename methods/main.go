package main

import "fmt"

//func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }

type S struct {
	a int
}

//new type
type SType S
type IntType int

//alias
type SAlias = S

func (recv S) print() {
	fmt.Println("Please print")
}

func (recv SType) print() {
	fmt.Println("Please print")
}

func (recv IntType) print() {
	fmt.Println("Please print int type", recv)
}

func main() {
	a := S{11}
	a.print()

	b := SType{10}
	b.print()

	c := IntType(3)
	c.print()
}
