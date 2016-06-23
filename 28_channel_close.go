package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)
	go sendMessage(messages)
	printMessage(messages)
}

func sendMessage(ch chan<- string) {
	for i := 0; i <= 9; i++ {
		ch <- fmt.Sprintln("data ", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for data := range ch {
		fmt.Println(data)
	}
}
