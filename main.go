package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ Калькулятор индекса массы тела (IMT) __")
	userWidth, userHeight := getUserInput()
	IMT := calculateIMT(userWidth, userHeight)
	outputResult(IMT)
}

func calculateIMT(userKg float64, userheight float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userheight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWidth float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWidth)

	return userWidth, userHeight
}

func outputResult(imt float64) {
	fmt.Printf("Ваш индекс массы тела: %.0f", imt)
}
