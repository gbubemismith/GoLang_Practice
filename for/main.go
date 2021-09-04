package main

import "fmt"

func main() {
	i := 0

	//while loop -- this is equivalent to while loop in other languages e.g C#, Java, Javascript etc
	for i < 5 {
		fmt.Println("Number", i)
		i++
	}

	str := "Test string to test for range"

	//for range -- this is equivale to for each in c#
	for index, val := range str {
		fmt.Printf("Charcter pos %d is: %c \n", index, val)
	}

}
