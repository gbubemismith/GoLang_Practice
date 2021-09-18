package main

func changeValue(str *string) {
	*str = "changed"
}

func changeValue2(str string) {
	str = "changed"
}

func main() {

	//1. & -> turn value to memory address
	//2. * -> turn memory address to value

	// x := 7
	// y := &x
	// *y = 8
	// fmt.Println(*y, x)

	toChange := "hello"
	changeValue(&toChange)
}
