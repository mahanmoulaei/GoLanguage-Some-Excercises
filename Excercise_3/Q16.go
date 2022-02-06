package main

import (
	"fmt"
	"math"
)

func main() {

	var inputDays float64
	var inputKwh float64

	fmt.Print("Enter the amount of days the electricity was used : ")
	fmt.Scan(&inputDays)
	fmt.Print("Enter the amount of kilowatt hours(kWh) consumed : ")
	fmt.Scan(&inputKwh)

	calcElecBill(inputDays, inputKwh)

}

func calcElecBill(days float64, kwh float64) {

	var origKwhRate float64 = .30
	var dayRate float64 = .50
	var newKwhRate float64 = .20
	var totalBill float64 = 0
	var extraKwh float64 = 0
	var notExtraKwh float64 = 0

	if days < 0 || kwh < 0 {
		fmt.Print("Please enter valid numbers\n")
		main()
	} else if kwh <= 200 {
		totalBill = math.Round((days*dayRate+kwh*origKwhRate)*100) / 100
		fmt.Printf("The total electricity bill is $%g", totalBill)
	} else {
		extraKwh = kwh - 200
		notExtraKwh = kwh - extraKwh
		totalBill = math.Round((notExtraKwh*origKwhRate+extraKwh*newKwhRate+days*dayRate)*100) / 100
		fmt.Printf("The total electricity bill is $%g", totalBill)
	}
}
