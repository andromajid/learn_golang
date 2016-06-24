package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var message = make(chan int)
	go sendData(message)
	retriveData(message)
}

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retriveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Println("Receive data : ", data)
		case <-time.After(time.Second * 5):
			fmt.Println("Timeout. no activities in 5 seconds")
			break loop
		}
	}
}
