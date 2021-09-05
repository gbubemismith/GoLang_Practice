package main

import "fmt"

func main() {
	//slice is like a list or an arraylist in c# and Java

	//difference between slice and array initialization
	//array has the length initialized e.g
	testArr := [3]int{1, 2, 3}
	_ = testArr
	//while a slic is declared without the length e.g
	testSlice := []int{1, 2, 3, 4, 5}
	_ = testSlice

	//initialization
	s := []int{1, 2, 3, 4}
	_ = s

	//make function can be used to create a slice
	a := make([]int, 5, 5)
	fmt.Println(a)
}

//passing a slice to a function
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
