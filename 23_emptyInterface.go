package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{}
	secret = "andro majid"
	fmt.Println(secret)

	secret = []string{"andro", "majid", "arkananta"}
	fmt.Println(secret)
	//interface kosong bisa jadi sebuah hash map
	var data map[string]interface{}
	data = map[string]interface{}{
		"name":      "Arka hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}
	fmt.Println(data)
	//type assertion pada interface
	secret = 12
	var number = secret.(int) * 10
	fmt.Println(number)
	secret = []string{"andro", "majid"}
	var namaPanjang = strings.Join(secret.([]string), " ")
	fmt.Println("nama panjang saya : ", namaPanjang)
}
