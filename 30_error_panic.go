package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Println("masukkan angka")
	fmt.Scanln(&input)

	var number int
	var err error

	number, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println("these is the ", number, " real number")
	} else {
		fmt.Println("these are not number ", input)
		fmt.Println(err.Error())
	}
	input = ""
	fmt.Println("masukkan karakter")
	fmt.Scanln(&input)
	if valid, err := validate(input); valid {
		fmt.Println("helllo", input)
	} else {
		fmt.Println(err.Error())
		//showing panic to trace error
		panic(err.Error())
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}
