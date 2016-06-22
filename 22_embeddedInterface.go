package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k kubus) luas() float64 {
	return math.Pow(k.sisi, 2)
}

func (k kubus) keliling() float64 {
	return k.sisi * 4
}

func (k kubus) volume() float64 {
	return k.luas() * k.sisi
}

func main() {
	var bangunRuang hitung = &kubus{3}
	fmt.Println("===== kubus")
	fmt.Println("luas      :", bangunRuang.luas())
	fmt.Println("keliling  :", bangunRuang.keliling())
	fmt.Println("volume    :", bangunRuang.volume())
}
