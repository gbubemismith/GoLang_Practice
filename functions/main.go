package main

import "fmt"

func main() {
	// fmt.Println(SumInts(1, 2, 3, 4, 5))
	callback(1, Add)

}

func example1(name string) string { //takes 1 parameter and returns a string
	return ""
}

func SumProductDiff(a, b int) (int, int, int) { //takes two parameters as integers and returns three results
	n1 := a + b
	n2 := a * b
	n3 := a - b

	return n1, n2, n3
}

func SumInts(list ...int) (sum int) {

	for _, v := range list {
		sum += v
	}

	return
}

func Factorial(n uint64) (fac uint64) {
	if n <= 1 {
		return 1
	}

	fac = n * Factorial(n-1)
	return
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}
