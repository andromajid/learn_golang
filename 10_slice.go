package main

import "fmt"

func main() {
	//perbedaan slice dan array adalah inisiasy
	var fruitArray = [4]string{
		"pisang",
		"rambutan",
		"nangka",
		"jambu",
	}
	/*var fruitSlice = []string{
		"pisang",
		"rambutan",
		"nangka",
		"jambu",
	}*/

	var newFruitSlice = fruitArray[2:3]
	fmt.Println(newFruitSlice[0])
	newFruitSlice[0] = "nangka Baru"
	fmt.Println(newFruitSlice[0])
	fmt.Println(fruitArray[2])
}
