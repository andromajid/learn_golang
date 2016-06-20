package main

import (
	"fmt"
	"strings"
)

func main() {
	var siswa1 = student{
		name:  "arkananta majid",
		grade: 12,
	}
	var name = siswa1.getNameAt(2)
	siswa1.sayHello()
	fmt.Println("nama : ", name)
}

type student struct {
	name  string
	grade int
}

func (siswa student) sayHello() {
	fmt.Println("hello nama saya : ", siswa.name)
}

func (siswa student) getNameAt(index int) string {
	return strings.Split(siswa.name, " ")[index-1]
}
