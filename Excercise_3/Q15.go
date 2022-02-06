package main

import (
	"fmt"
	"math"
)

func main() {

	var inputBalance float64
	var inputWithdrawal float64

	fmt.Print("Enter your balance : ")
	fmt.Scan(&inputBalance)
	fmt.Print("Enter amount to withdrawal : ")
	fmt.Scan(&inputWithdrawal)

	atm(inputBalance, inputWithdrawal)

}

func atm(balance float64, withdrawal float64) {

	var newBalance float64 = 0
	if withdrawal > balance {
		fmt.Print("Error! Insufficient funds. Please try again\n")
		main()
	} else {
		newBalance = math.Round((balance-withdrawal)*100) / 100
		fmt.Printf("Your new balance is $%g", newBalance)
	}
}
