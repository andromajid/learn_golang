package main

import (
	"fmt"
)

func main() {
	var getMinMax = func(index []int) (int, int) {
		var min, max int
		for i, val := range index {
			switch {
			case i == 0:
				max, min = val, val
			case val > max:
				max = val
			case val < min:
				min = val
			}
		}
		return min, max
	}
	var number = []int{1, 2, 3, 44, 12, 77, 2}
	var min, max = getMinMax(number)
	fmt.Println("Nilai Minimal : ", min)
	fmt.Println("nilai maximal : ", max)

	//closure dengan iife (Immediately-Invoked Function Expression)
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var newNumber = func(min int) []int {
		var newArray []int
		for _, data := range numbers {
			if data < min {
				continue
			}
			newArray = append(newArray, data)
		}
		return newArray
	}(3)
	fmt.Println("data awal :", numbers)
	fmt.Println("data setelah sorting : ", newNumber)

	var total, clos = filterMax(numbers, 2)
	fmt.Println("total Data : ", total)
	fmt.Println("closure", clos())
}

//closure sebagai nilai kembalian
func filterMax(number []int, max int) (int, func() []int) {
	var returnData []int
	for _, data := range number {
		if data <= max {
			returnData = append(returnData, data)
		}
	}
	return len(returnData), func() []int {
		return returnData
	}
}
