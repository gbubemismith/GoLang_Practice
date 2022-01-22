package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool   "An important answer"
	field2 string `the name of one thing`
	field3 int    "How much there are"
}

func main() {
	tt := TagType{true, "Smith", 1}

	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	fmt.Println("ttType", ttType)
	ixField := ttType.Field(ix)
	fmt.Println("ixField", ixField.Tag)
}
