package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"] = 1
	chicken["february"] = 2
	fmt.Println("januari ", chicken["januari"])
	fmt.Println("februari ", chicken["february"])
	//menginisasi pada saat awal map
	var chickenAwal = map[string]int{
		"januari":  10,
		"februari": 20,
	}
	fmt.Println("januari ", chickenAwal["januari"])
	fmt.Println("februari ", chickenAwal["februari"])
	//iterasi dengan map
	for key, val := range chickenAwal {
		fmt.Println(key, " : ", val)
	}

	fmt.Println("menghapus sebuah map")
	delete(chickenAwal, "januari")
	fmt.Println("total data sekarang : ", len(chickenAwal))
}
