package main

import (
	"fmt"
)

func main() {
	var day1, month1, year1, day2, month2, year2 int

	for {
		fmt.Println("Enter a date (DD/MM/YYYY): ")
		fmt.Scanf("%d/%d/%d\n", &day1, &month1, &year1)

		if !isValidDate(day1, month1, year1) {
			fmt.Println("The entered date is invalid!")
			continue
		} else {
			break
		}
	}

	for {
		fmt.Println("Enter another date (DD/MM/YYYY): ")
		fmt.Scanf("%d/%d/%d\n", &day2, &month2, &year2)

		if !isValidDate(day2, month2, year2) {
			fmt.Println("The entered date is invalid!")
			continue
		} else {
			break
		}
	}

	if year1 < year2 {
		fmt.Printf("The first entered date (%d/%d/%d) comes before the second entered date (%d/%d/%d)", day1, month1, year1, day2, month2, year2)
	} else if year2 < year1 {
		fmt.Printf("The second entered date (%d/%d/%d) comes before the first entered date (%d/%d/%d)", day2, month2, year2, day1, month1, year1)
	} else {
		if month1 < month2 {
			fmt.Printf("The first entered date (%d/%d/%d) comes before the second entered date (%d/%d/%d)", day1, month1, year1, day2, month2, year2)
		} else if month2 < month1 {
			fmt.Printf("The second entered date (%d/%d/%d) comes before the first entered date (%d/%d/%d)", day2, month2, year2, day1, month1, year1)
		} else {
			if day1 < day2 {
				fmt.Printf("The first entered date (%d/%d/%d) comes before the second entered date (%d/%d/%d)", day1, month1, year1, day2, month2, year2)
			} else if day2 < day1 {
				fmt.Printf("The second entered date (%d/%d/%d) comes before the first entered date (%d/%d/%d)", day2, month2, year2, day1, month1, year1)
			} else {
				fmt.Printf("The entered dates are equal...")
			}
		}
	}

}

func isValidDate(day int, mon int, year int) bool {
	var is_valid bool = true
	var is_leap bool = false

	if year >= 1800 && year <= 9999 {
		//  check whether year is a leap year
		if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
			is_leap = true
		}

		// check whether mon is between 1 and 12
		if mon >= 1 && mon <= 12 {
			// check for days in feb
			if mon == 2 {
				if is_leap && (day == 29) {
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
