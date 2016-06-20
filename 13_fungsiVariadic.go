package main

import (
	"fmt"
)

func main() {
	var rata2 = calculate(1, 2, 4, 5, 6, 77, 11)
	fmt.Println("rata2 : ", rata2)
}

/*
fungsi di bawah adalah fungsi varidic
*/
func calculate(number ...int) float64 {
	var totalInt int = 0
	for _, val := range number {
		totalInt += val
	}
	var avg = float64(totalInt) / float64(len(number))
	return avg
}
