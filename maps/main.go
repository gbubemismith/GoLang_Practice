package main

import "fmt"

func main() {
	// var isPresent bool

	//var map1 map[keytype]valuetype
	var mapLit map[string]int //nil map
	mapLit = map[string]int{"one": 1, "two": 2}
	fmt.Println(mapLit)        //print map
	fmt.Println(mapLit["one"]) //get value of key

	m := make(map[string]int) //make to create a map
	m["route"] = 66           //initializing a key with value

	// _, isPresent = m["route"]
	// fmt.Println(isPresent)
	if _, isPresent := m["route"]; isPresent { //
		fmt.Println("Yes")
	}

	delete(m, "route")
	if _, isPresent := m["route"]; isPresent { //
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
