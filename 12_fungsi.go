package main

import "fmt"
import (
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var names = []string{"andro", "majid"}
	printMessage("Namaku : ", names)
	//create random with value
	rand.Seed(time.Now().Unix())
	var randomValue int
	randomValue = randomWithValue(2, 10)
	fmt.Println("random value 1 :", randomValue)
	randomValue = randomWithValue(2, 10)
	fmt.Println("random value 2 :", randomValue)
	randomValue = randomWithValue(2, 10)
	fmt.Println("random value 3 :", randomValue)

	//fungsi dengan multiple return
	var area, keliling = calculate(10)
	fmt.Println("luas lingkaran dengan diameter 10 : ", area)
	fmt.Println("keliling lingkaran dengan diagram 10 : ", keliling)
}

func printMessage(message string, array []string) {
	var namesString = strings.Join(array, " ") //seperti php implode
	fmt.Println(message, namesString)
}

func randomWithValue(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func calculate(jari2 float64) (float64, float64) {
	//luas
	var area = math.Pi * math.Pow(jari2/2, 2)
	//keliling
	var keliling = math.Pi * jari2

	return area, keliling
}
