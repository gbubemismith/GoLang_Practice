package main

import "fmt"

func main() {
	//slice is like a list or an arraylist in c# and Java

	//difference between slice and array initialization
	//array has the length initialized e.g
	// testArr := [3]int{1, 2, 3}
	// _ = testArr
	// //while a slice is declared without the length e.g
	// testSlice := []int{1, 2, 3, 4, 5}
	// _ = testSlice

	// //initialization
	// s := []int{1, 2, 3, 4}
	// _ = s

	// //make function can be used to create a slice
	// a := make([]int, 5, 5)
	// fmt.Println(a)

	//-------------------reslicing
	// slice1 := make([]int, 0, 10)
	// fmt.Println(slice1)

	// slice1 = slice1[0 : len(slice1)+1] //extend slice1 length by 1
	// fmt.Println(len(slice1))

	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7] // starts at index 5 end ends before index 7
	fmt.Println(cap(a))
	a = a[0:4]
	fmt.Println(a)

	//append() function is used to add more elements
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5, 6)

	//make() function is used to preallocate a slice
	//e.g slice := make([]int, 10)

}

//passing a slice to a function
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
