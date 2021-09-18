package main

import "fmt"

func main() {
	name := new(string)

	fmt.Println(*name)

	//new returns a pointer
	//make is used in maps, slices, channels
}
