package main

import (
	"fmt"
	"math"
)

func main() {

	var inputFoodAndHouse float64
	var inputCommonExp float64
	var inputPublicTrans float64
	var inputRent float64
	var inputMonthlyBill float64
	var inputPaycheque float64

	fmt.Print("Enter your food and house expenses for the week : ")
	fmt.Scan(&inputFoodAndHouse)
	fmt.Print("Enter your common expenses for the week : ")
	fmt.Scan(&inputCommonExp)
	fmt.Print("Enter your public transport expenses for the month : ")
	fmt.Scan(&inputPublicTrans)
	fmt.Print("Enter your rent for the month : ")
	fmt.Scan(&inputRent)
	fmt.Print("Enter your total of monthly bills : ")
	fmt.Scan(&inputMonthlyBill)
	fmt.Print("Enter 1 of your average paycheques : ")
	fmt.Scan(&inputPaycheque)

	CalculateExpenses(inputFoodAndHouse, inputCommonExp, inputPublicTrans, inputRent, inputMonthlyBill, inputPaycheque)

}

func CalculateExpenses(foodAndHouse float64, commonExp float64, publicTrans float64, rent float64, monthlyBill float64, paycheque float64) {

	var weeklyExp float64 = math.Round((foodAndHouse+commonExp)*4*100) / 100
	var monthlyExp float64 = math.Round(publicTrans+rent+monthlyBill) * 100 / 100
	var totalExp float64 = weeklyExp + monthlyExp
	var totalPay float64 = math.Round(paycheque*2*100) / 100
	var difference float64 = math.Round((totalPay-totalExp)*100) / 100

	if foodAndHouse < 0 || commonExp < 0 || publicTrans < 0 || publicTrans < 0 || rent < 0 || monthlyBill < 0 {
		fmt.Print("Please enter positive numbers!\n")
		main()
	} else {
		fmt.Printf("Your total expenses for the month are : $%g\nYour total income for the month is $%g\nYou have $%g left to spend this month", totalExp, totalPay, difference)
	}
}
