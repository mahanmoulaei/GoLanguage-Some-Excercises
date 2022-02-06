package main

import (
	"fmt"
	"math"
)

func main() {

	var inputPrice float64
	var inputQuantity int

	fmt.Print("Enter a price : ")
	fmt.Scan(&inputPrice)
	fmt.Print("Enter the quantity : ")
	fmt.Scan(&inputQuantity)

	CalculateTax(inputPrice, inputQuantity)

}

func CalculateTax(price float64, quantity int) {

	var totalPrice = price * float64(quantity)
	var gst float64 = math.Round(totalPrice*.07*100) / 100
	var qst float64 = math.Round(gst+totalPrice) * .075 * 100 / 100
	var finalPrice float64 = gst + qst + math.Round(totalPrice*100)/100

	fmt.Printf("The GST is : %g\nThe QST is : %g\nThe total price is %g", gst, qst, finalPrice)
}
