package main

import (
	"belajar-golang/accessLevel/library"
	"fmt"
)

func main() {
	fmt.Printf("Name  : %s\n", library.Student.Name)
	fmt.Printf("Grade : %d\n", library.Student.Grade)
}
