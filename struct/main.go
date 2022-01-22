package main

import (
	"fmt"
	"unsafe"
)

//structs are like the equivalent of classes in OOP languages
/**
general format of defining a struct

	type identifier struct {
		field1 type1
		field2 type2
		...
	  }
**/

type myStruct struct {
	i int
	j float32
	k string
}

//implementign struct factory in Go
type File struct {
	fd   int
	name string
}

//factory (equivalent of a constructor in OOP languages)
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func main() {
	// var v *myStruct
	//struct literal
	// check := myStruct{10, 5.5, "name"}
	// check.i

	//intializing a struct
	v := myStruct{10, 5.5, "name"}
	_ = v
	//or
	var mt myStruct
	mt = myStruct{10, 5.5, "name"}
	_ = mt
	//or
	// ms := new(myStruct) //returns a pointer
	// ms.i

	//get size of struct
	fmt.Println("Size of struct", unsafe.Sizeof(myStruct{}))

	//anonymous struct
	sample := struct {
	}{}

	_ = sample

}
