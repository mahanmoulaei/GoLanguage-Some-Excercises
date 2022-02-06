package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	var salaryString string
	var salaryFloat float64
	var isMarried bool
	var noOfChildInt int
	var isNewcomer bool

	for { //Salary Amount

		input = ""
		salaryString = ""
		salaryFloat = 0.0
		isMarried = false
		noOfChildInt = 0
		isNewcomer = false

		fmt.Println("*** TAX CALCULATION ***")
		fmt.Println("Enter your annual salary : (type exit to quit the program)")
		fmt.Scan(&input)

		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				salaryString = fmt.Sprintf("%.2f", float)
				salaryFloat, error = strconv.ParseFloat(salaryString, 64)

				for { //Marital Status

					fmt.Println("Are you Married? (Y/N) : (type exit to quit the program)")
					fmt.Scan(&input)

					if input == "exit" {
						os.Exit(0)
					} else {
						input = strings.ToLower(input)
						if input == "y" || input == "yes" {
							isMarried = true
						} else if input == "n" || input == "no" {
							isMarried = false
						} else {
							continue
						}

						for { //Number of Children(s)
							fmt.Println("How many children do you have? : (type exit to quit the program)")
							fmt.Scan(&input)

							if input == "exit" {
								os.Exit(0)
							} else {
								if noOfChildInt, error = strconv.Atoi(input); error == nil {

									for { //Newcomer to the province

										fmt.Println("Are you a newcomer to the province? (Y/N) : (type exit to quit the program)")
										fmt.Scan(&input)

										if input == "exit" {
											os.Exit(0)
										} else {
											input = strings.ToLower(input)
											if input == "y" || input == "yes" {
												isNewcomer = true
											} else if input == "n" || input == "no" {
												isNewcomer = false
											} else {
												continue
											}

											tax := CalculateTax(salaryFloat, isMarried, noOfChildInt, isNewcomer)
											if tax <= 0 {
												fmt.Println("You DO NOT need to pay any tax! Be Happy...")
											} else {
												taxAmount := (salaryFloat * tax) / 100
												fmt.Println("The amount of tax is you need to pay is %", tax, " of your yearly income which is $", taxAmount)
											}
											break
										}
									}
								} else {
									continue
								}
								break
							}
						}
						break
					}
				}
			}
		}
	}
}

func CalculateTax(salaryFloat float64, isMarried bool, noOfChildInt int, isNewcomer bool) float64 {
	var tax float64 = 0

	if salaryFloat > 0.0 && salaryFloat <= 18000.0 {
		tax = 10 + CalculateTaxReduction(isMarried, noOfChildInt, isNewcomer)
	} else if salaryFloat > 18000.0 && salaryFloat <= 32000.0 {
		tax = 20 + CalculateTaxReduction(isMarried, noOfChildInt, isNewcomer)
	} else if salaryFloat > 32000.0 && salaryFloat <= 60000.0 {
		tax = 30 + CalculateTaxReduction(isMarried, noOfChildInt, isNewcomer)
	} else if salaryFloat > 60000.0 {
		tax = 40 + CalculateTaxReduction(isMarried, noOfChildInt, isNewcomer)
	} else {
		fmt.Println("Run the program again and check your entry values...")
		os.Exit(0)
	}

	return tax
}

func CalculateTaxReduction(isMarried bool, noOfChildInt int, isNewcomer bool) float64 {
	var taxReduction float64 = 0

	if isMarried {
		taxReduction -= 2
	}
	if noOfChildInt > 0 {
		for i := 1; i <= noOfChildInt; i++ {
			taxReduction -= 0.5
		}
	}
	if isNewcomer {
		taxReduction -= 8
	}

	return taxReduction
}
