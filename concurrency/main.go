package main

import (
	"fmt"
	"time"
	"unicode"
)

//Defer -> postpone an operation till after function completes (clean up resources or execute after function finishes)
//Defer puts the deferred operation in a stack and executes in reverse order

//Concurrency -> allows

//Goroutines -> allows functions to run concurrently

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune

	//closure
	capIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
		fmt.Printf("%c done!\n", r)
	}

	for i := 0; i < len(data); i++ {
		go capIt(data[i])
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Capitalized: %c\n", capitalized)
}
