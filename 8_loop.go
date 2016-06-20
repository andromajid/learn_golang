package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println("angka ", i)
	}

	//looping seperti while
	loop := true
	i := 0
	for loop {
		fmt.Println("angka2 ", i)
		if i == 5 {
			loop = false
		}
		i++
	}
}
