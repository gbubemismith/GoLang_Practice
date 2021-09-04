package main

func callback(y int, f func(int, int)) {
	f(y, 2)
}
