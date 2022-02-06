package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	for {
		fmt.Println("Enter a year to check if it's a leap year : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				if integer >= 1000 && integer <= 9999 {
					if isValidYear(integer) {
						fmt.Printf("Year %d IS a leap year!\n", integer)
					} else {
						fmt.Printf("Year %d is NOT a leap year!\n", integer)
					}
				} else {
					fmt.Println("Enter a valid year between 1000 and 9999!")
				}
			}
		}
		continue
	}
}

func isValidYear(year int) bool {
	var is_leap bool = false

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		is_leap = true
	} else {
		is_leap = false
	}

	return is_leap
}
