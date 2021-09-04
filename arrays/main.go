package main

import "fmt"

func main() {

	//different ways to initialize arrays
	// just size of the array
	var arr [5]int

	//size and values initialized
	var arrAge = [5]int{1, 2, 3, 4, 5}
	_ = arrAge

	//values determine the size
	var arrAge2 = []int{1, 2, 3, 4, 5}
	_ = arrAge2

	//compiler determins the size
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	_ = arrLazy

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	fmt.Println(arr)

	//pointers with arrays
	array := [3]float64{1.0, 2.0, 3.0}
	x := Sum(&array)
	fmt.Println(x)

}

//receives a pointer
func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}
