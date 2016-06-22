package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var message = make(chan string)
	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello : %s", who)
		message <- data
	}
	go sayHelloTo("andro")
	go sayHelloTo("arka")
	go sayHelloTo("jason bourne")

	var message1 = <-message
	fmt.Println(message1)

	var message2 = <-message
	fmt.Println(message2)

	var message3 = <-message
	fmt.Println(message3)

	var newMessage = make(chan string)
	for _, data := range []string{"andro", "majid", "alatas"} {
		go func(who string) {
			var data = fmt.Sprintf("hello goRoutine %s", who)
			newMessage <- data
		}(data)
	}
	for i := 0; i < 3; i++ {
		printMessage(newMessage)
	}
}

func printMessage(what chan string) {
	fmt.Println(<-what)
}
