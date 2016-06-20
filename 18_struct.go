package main

import (
	"fmt"
)

func main() {
	var siswa1 student
	siswa1.name = "arkananta majid"
	siswa1.grade = 2
	fmt.Println("namaku adalah : ", siswa1.name)
	fmt.Println("aku kelas : ", siswa1.grade)
	//inisiasi dengan method lain
	var siswa2 = student{"Andro Majid", 2}
	var siswa3 = student{
		name:  "arka majid",
		grade: 2,
	}
	fmt.Println("Siswa Nama", siswa2.name)
	fmt.Println("siswa 3 ", siswa3.name)
	//menggunakan embedded ala inheritence
	siswaInheretence := studentInheritance{}
	siswaInheretence.age = 4
	siswaInheretence.grade = 1
	siswaInheretence.name = "andro"
	fmt.Printf(`
	nama : %s
	kelas : %i
	umur : %i
	`, siswaInheretence.name, siswaInheretence.grade, siswaInheretence.age)

	//using array and anonymous struct
	var allStudent = []struct {
		person
		grade int
	}{
		{person: person{"andro", 2}, grade: 2},
		{person: person{"arka", 2}, grade: 3},
	}

	for _, student := range allStudent {
		fmt.Println(student.name)
	}
}

type studentInheritance struct {
	grade int
	person
}

type student struct {
	name  string
	grade int
}

type person struct {
	name string
	age  int
}

//below is nested struct

type studentData struct {
	person struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}
