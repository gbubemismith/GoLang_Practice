package main

import "fmt"

func main() {
	// s := []int{1, 2, 3}
	// fmt.Println("Capacity before", cap(s))
	// s = enlarge(s, 5)

	// fmt.Println(s)
	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := insertSlice(s, in, 0)
	fmt.Println(res)

}

func enlarge(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	fmt.Println("Capacity in method", cap(ns), len(ns))
	copy(ns, s)
	return ns
}

func insertSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index]) // copying all elements from 0 index till the end
	fmt.Println("at::", at)
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}
