package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("these would be executed after the program end")
	fmt.Println("hello")
	defer fmt.Println("execute 2 defer")
}
