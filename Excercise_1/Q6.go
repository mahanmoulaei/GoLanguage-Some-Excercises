package main

import (
	"fmt"
	"math"
)

func main() {
	var inputNumVehicules int

	fmt.Print("How many vehicules have you sold? : ")
	fmt.Scan(&inputNumVehicules)

	calcCompensation(inputNumVehicules)
}

func calcCompensation(vehicules int) {

	var comission float64 = math.Round(float64(vehicules)*200*100) / 100
	var bonus float64 = math.Round(comission*.05*100) / 100
	var totalComp float64 = comission + bonus
	var totalEarnings float64 = comission + bonus + 400

	if vehicules < 0 {
		fmt.Print("Please enter a positive number\n")
		main()
	} else {
		fmt.Printf("The total amount of compensation without base salary is $%g\nThe total amount including base salary is $%g", totalComp, totalEarnings)
	}
}
