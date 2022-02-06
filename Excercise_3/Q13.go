package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	var inputQuantity int
	var inputAnswer string

	fmt.Print("Enter quantity of diskettes : ")
	fmt.Scan(&inputQuantity)
	fmt.Print("Are you a member of Club Z? (Y/N) : ")
	fmt.Scan(&inputAnswer)

	calcDiskPrice(inputQuantity, inputAnswer)

}

func calcDiskPrice(quantity int, answer string) {

	var answerUpper string = strings.ToUpper(answer)
	var origPrice float64 = 1
	var discountPrice float64 = .70
	var totalPrice float64 = 0
	var member float64 = 0
	var finalPrice float64 = 0

	if quantity < 0 {
		main()
	} else if quantity < 25 && answerUpper == "N" {
		finalPrice = math.Round((float64(quantity)*origPrice)*100) / 100
		fmt.Printf("The total Price is $%g", finalPrice)
	} else if quantity >= 25 && answerUpper == "Y" {
		totalPrice = float64(quantity) * discountPrice
		member = totalPrice * 0.02
		finalPrice = math.Round((totalPrice-member)*100) / 100
		fmt.Printf("The total price is $%g", finalPrice)
	} else if quantity < 25 && answerUpper == "Y" {
		totalPrice = float64(quantity) * origPrice
		member = totalPrice * 0.02
		finalPrice = math.Round((totalPrice-member)*100) / 100
		fmt.Printf("The total price is $%g", finalPrice)
	} else if quantity >= 25 && answerUpper == "N" {
		finalPrice = math.Round((float64(quantity)*discountPrice)*100) / 100
		fmt.Printf("The total price is $%g", finalPrice)
	}

}
