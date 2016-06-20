package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"andro", "majid", "arkananta"}
	var dataString = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	fmt.Println("data di dalam adalah : ", dataString)
	var dataInt = filter(data, func(each string) bool {
		return len(each) == 5
	})
	fmt.Println("jumlah kata yang ada lima huruf :", dataInt)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
