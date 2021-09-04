package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	fplus := func(x, y int) int { return x + y }

	fmt.Println(fplus(3, 4))

	//invovked directly
	func(x, y int) int {
		return x + y
	}(3, 4)

	//debugging with closures
	where := func() { // debugging function
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	// some code
	a := 2 * 5
	where()
	// some more code
	b := a + 100
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	where()
}
