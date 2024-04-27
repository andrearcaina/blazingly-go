package main

import (
	"fmt"
)

func main() {
	var ch = make(chan int)
	var c = make(chan int)
	//ch <- 1 like c : [1] this will block until the value is read (and the channel only has a buffer of 1)
	// fmt.Println(<-ch) this creates a deadlock because the channel is empty and there is no goroutine to write to it

	go process(ch)
	go process(c)

	for i := range ch {
		fmt.Println(i)
	}

	// could also do:
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func process(ch chan int) {
	defer close(ch) // this will close the channel when the function returns
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
