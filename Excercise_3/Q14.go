package main

import (
	"fmt"
)

func main() {

	var inputCopies float64
	fmt.Print("Enter the amount of copies : ")
	fmt.Scan(&inputCopies)

	priceOfCopies(inputCopies)

}

func priceOfCopies(copies float64) {

	var totalPrice float64 = 0
	var pricePerCopy float64 = 0
	var newPricePerCopy float64 = 0

	if copies < 100 {
		pricePerCopy = .05
		totalPrice = copies * pricePerCopy
		fmt.Printf("The total price of copies is $%g", totalPrice)
	} else if copies > 100 {
		var extraCopies float64 = copies - 100
		var newTotal float64 = copies - extraCopies
		newPricePerCopy = .05
		pricePerCopy = .03
		totalPrice = (newTotal * pricePerCopy) + (newPricePerCopy * extraCopies)
		fmt.Printf("The total price of copies is $%g", totalPrice)
	}
}
