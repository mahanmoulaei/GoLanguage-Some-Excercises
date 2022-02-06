package main

import (
	"fmt"
	"math"
)

func main() {

	var inputPayed float64
	fmt.Print("Enter amount payed : ")
	fmt.Scan(&inputPayed)

	calcTip(inputPayed)

}

func calcTip(amountPayed float64) {
	if amountPayed >= 1 {
		var tip float64 = math.Round(amountPayed*.15*100) / 100
		fmt.Printf("Your tip is : %g", tip)
	} else {
		var tip float64 = 0
		fmt.Printf("Your tip is : %g", tip)
	}

}
