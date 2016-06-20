package main

import "fmt"

func main() {
	point := 2
	if point == 1 {
		fmt.Println("ini adalah satu")
	} else if point == 2 {
		fmt.Println("ini adalah dua")
	} else {
		fmt.Println("ini bukan angka")
	}

	nilai := 8840.0

	if percent := nilai / 100; percent > 100 {
		fmt.Printf("%.1f%s bener", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good", percent, "%")
	} else {
		fmt.Printf("%.1f%s not good enough", percent, "%")
	}
}
