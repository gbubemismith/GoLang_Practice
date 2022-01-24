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
	// * with a data type inidcates that value is a pointer e.g var valuePtr *int, valuePtr = &val

	//dereferencing-> * with a pointer varaible gives the actual value
	//e.g val = *valuePtr 

	// x := 7
	// y := &x
	// *y = 8
	// fmt.Println(*y, x)

	toChange := "hello"
	changeValue(&toChange)
}
