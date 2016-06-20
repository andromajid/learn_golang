package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "andro"
	names[1] = "majid"
	names[2] = "arkananta"
	names[3] = "santyani"
	var total = len(names)
	for i := 0; i < total; i++ {
		fmt.Println("array ke ", i, " data : ", names[i])
	}
	//inisiasi array pada saat deklarasi
	var arrayFruit = [...]string{
		"andro",
		"majid",
		"santyani",
		"majid",
	}
	total = len(arrayFruit)
	for i := 0; i < total; i++ {
		fmt.Println("array ke ", i, " data : ", arrayFruit[i])
	}

	//looping dengan data arrayFruit seperti foreach
	for i, fruit := range arrayFruit {
		fmt.Println("elemen ke", i, " data : ", fruit)
	}

	//tidak butuh data key
	//looping dengan data arrayFruit seperti foreach dengan menggunakan variable blackhole
	for _, fruit := range arrayFruit {
		fmt.Println("data : ", fruit)
	}
}
