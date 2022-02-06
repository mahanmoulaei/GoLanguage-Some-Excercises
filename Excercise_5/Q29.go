package main

import (
	"fmt"
)

func main() {
	var day, month, year int

	for {
		fmt.Println("Enter a date (DD/MM/YYYY): ")
		fmt.Scanf("%d/%d/%d\n", &day, &month, &year)

		if !isValidDate(day, month, year) {
			fmt.Println("The entered date is invalid!")
			continue
		} else {
			break
		}
	}

	if isValidDate(day+1, month, year) {
		fmt.Printf("Today's date: %d/%d/%d => Tomorrow's date: %d/%d/%d\n", day, month, year, day+1, month, year)
	} else if isValidDate(1, month+1, year) {
		fmt.Printf("Today's date: %d/%d/%d => Tomorrow's date: %d/%d/%d\n", day, month, year, 1, month+1, year)
	} else if isValidDate(1, 1, year+1) {
		fmt.Printf("Today's date: %d/%d/%d => Tomorrow's date: %d/%d/%d\n", day, month, year, 1, 1, year+1)
	}

}

func isValidDate(day int, mon int, year int) bool {
	var is_valid bool = true

	if year >= 1800 && year <= 9999 {

		// check whether mon is between 1 and 12
		if mon >= 1 && mon <= 12 {
			// check for days in feb
			if mon == 2 {
				if day == 29 {
					is_valid = true
				} else if day > 28 {
					is_valid = false
				}
			} else if mon == 4 || mon == 6 || mon == 9 || mon == 11 { // check for days in April, June, September and November
				if day > 30 {
					is_valid = false
				}
			} else if day > 31 { // check for days in rest of the months i.e Jan, Mar, May, July, Aug, Oct, Dec
				is_valid = false
			}
		} else {
			is_valid = false
			//Print("Month is invalid!", mon)
		}
	} else {
		is_valid = false
		//Print("Year is invalid!", year)
	}

	return is_valid
}

func Print(any ...interface{}) {
	fmt.Println(any...)
}
