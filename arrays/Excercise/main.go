package main

import "fmt"

func main() {
	//loop to fill an array form 0 - 14
	var arr [15]int

	fmt.Println("before", arr)
	Fill014(&arr)
	fmt.Println("after", arr)
}

func Fill014(a *[15]int) {
	for i := 0; i <= 14; i++ {
		a[i] = i
	}
}
