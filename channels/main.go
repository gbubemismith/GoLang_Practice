package main

import (
	"fmt"
	"time"
)

//Channel -> offer a way of communication for go routines // write data at one end and read data out the other

func main() {
	//creating a channel
	channel := make(chan string)
	go process("order", channel)
	for {
		msg, open := <-channel
		if !open {
			break
		}
		fmt.Println(msg)
	}

	//

}

func process(item string, out chan string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second / 2)
		out <- item
	}
	close(out)
}
