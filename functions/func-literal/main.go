package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, op func(lhs, rhs int) int) int {
	return op(lhs, rhs)
}

func main() {
	fmt.Println("2 + 2 = ", compute(2, 2, add))

	fmt.Println("10 - 2 = ", compute(10, 2, func(lhs, rhs int) int {
		return lhs - rhs
	}))
}
