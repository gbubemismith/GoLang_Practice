package main

import "fmt"

func main() {
	values := [][]int{}

	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}

	values = append(values, row1)
	values = append(values, row2)

	//display first row
	fmt.Println(values[0])

	//display second row
	fmt.Println(values[1])

	//get first element
	fmt.Println(values[0][0])

	//print entire slice
	fmt.Println(values)

}
