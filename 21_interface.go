package main

import (
	"fmt"
	"math"
)

func main() {
	var bangunDatar hitung
	bangunDatar = persegi{21}
	fmt.Println("luas persegi : ", bangunDatar.luas())
	fmt.Println("Keliling persegi : ", bangunDatar.keliling())

	//now the lingkaran who has the same as hitung interface
	bangunDatar = lingkaran{7}
	fmt.Println("luas Lingkaran : ", bangunDatar.luas())
	fmt.Println("Keliling Lingkaran : ", bangunDatar.keliling())
	fmt.Println("jari2 nya :", bangunDatar.(lingkaran).jariJari())
}

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (ling lingkaran) jariJari() float64 {
	return ling.diameter / 2
}

func (ling lingkaran) luas() float64 {
	return math.Pi * math.Pow(ling.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}
