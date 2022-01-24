package main

import "fmt"

//Variadic is represented with 3 dots e.g ...int -> means we can have any number of int values

func sum(nums ...int) int {
	fmt.Println(nums)
	return 0
}

func main() {
	a := []int{1, 2, 3}

	ans := sum(a...)
	fmt.Println(ans)
}
