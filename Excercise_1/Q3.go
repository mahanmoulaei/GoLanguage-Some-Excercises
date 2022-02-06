package main

import (
	"fmt"
	"math"
)

func main() {

	var inputHourlyRate float64
	var inputHoursWorked float64

	fmt.Print("Enter your hourly rate : ")
	fmt.Scan(&inputHourlyRate)
	fmt.Print("Enter your hours worked : ")
	fmt.Scan(&inputHoursWorked)

	CalcGrossSal(inputHourlyRate, inputHoursWorked)

}

func CalcGrossSal(hourlyRate float64, hoursWorked float64) {

	var weeklySalary float64 = math.Round(hourlyRate*hoursWorked*100) / 100

	if hourlyRate <= 0 {
		fmt.Print("Please enter a valid hourly rate : ")
		fmt.Scan(&hourlyRate)
		CalcGrossSal(hourlyRate, hoursWorked)
	} else if hoursWorked > 168 || hoursWorked <= 0 {
		fmt.Print("Please enter a valid amount of hours : ")
		fmt.Scan(&hoursWorked)
		var weeklySalary float64 = math.Round(hourlyRate*hoursWorked*100) / 100
		fmt.Printf("Your weekly salary is %g", weeklySalary)

	} else if hoursWorked <= 168 && hoursWorked > 0 {
		fmt.Printf("Your weekly salary is %g", weeklySalary)
	}

}
