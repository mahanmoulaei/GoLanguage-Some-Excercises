package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	var currentPremiumPrice float64
	var numberOfAccidents int
	var newPremiumPrice float64

	for {
		fmt.Println("Enter your current premium price: (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if float > 0 { //check for negatives and 0
					currentPremiumPrice = float
					fmt.Println("Enter the number of accidents you had during the last year: (type exit to quit the program)")
					fmt.Scan(&input)
					if input == "exit" {
						os.Exit(0)
					} else {
						if number, err := strconv.Atoi(input); err == nil {
							numberOfAccidents = number
							break
						}
					}
				}
			}
		}
		continue
	}

	if numberOfAccidents == 0 {
		newPremiumPrice = ((2 + 100) * currentPremiumPrice) / 100
		fmt.Println("Your new premium price is $", newPremiumPrice, "which has a %2 increase in compare to your old premium price and based on the number of accidents you've had!")
	} else if numberOfAccidents == 1 || numberOfAccidents == 2 {
		newPremiumPrice = ((5 + 100) * currentPremiumPrice) / 100
		fmt.Println("Your new premium price is $", newPremiumPrice, "which has a %5 increase in compare to your old premium price and based on the number of accidents you've had!")
	} else if numberOfAccidents == 3 {
		newPremiumPrice = ((10 + 100) * currentPremiumPrice) / 100
		fmt.Println("Your new premium price is $", newPremiumPrice, "which has a %10 increase in compare to your old premium price and based on the number of accidents you've had!")
	} else if numberOfAccidents >= 4 {
		newPremiumPrice = ((30 + 100) * currentPremiumPrice) / 100
		fmt.Println("Your new premium price is $", newPremiumPrice, "which has a %30 increase in compare to your old premium price and based on the number of accidents you've had!")
	}
}
