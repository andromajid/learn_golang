package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)
	fmt.Println("tipe variable adalah ", reflectValue.Type())
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable adalah : ", reflectValue.Int())
	}
	//using interface so doesnt have to guest variable type
	fmt.Println("data interface adalah : ", reflectValue.Interface())

	//reflection using class
	var student1 = &student{"andro", 2}
	student1.getPropertyInfo()
}

type student struct {
	Name  string
	Grade int
}

func (s student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}
