package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Enter a binary number composed of 0s and 1s:")
	fmt.Scanln(&number)

	fmt.Println(getBinaryToDecimalValue(number))
}

func getBinaryToDecimalValue(input int) int {
	decimalValue := 0
	baseOf2 := 1

	var num int = input

	for num > 0 {
		var lastDigit int
		lastDigit = num % 10
		num = num / 10
		decimalValue += lastDigit * baseOf2
		baseOf2 = baseOf2 * 2
	}

	return decimalValue
}
