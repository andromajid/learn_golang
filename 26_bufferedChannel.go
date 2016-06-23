package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan int, 2)

	go func() {
		for {
			i := <-message
			fmt.Println("receive data", i)
		}
	}()
	for i := 0; i < 6; i++ {
		fmt.Println("send data", i)
		message <- i
	}
}
