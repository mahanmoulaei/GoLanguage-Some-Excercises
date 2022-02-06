package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var input string
	fmt.Println("*** Grade Calculator ***")
	for {
		fmt.Println("Enter a grade between 0 and 100 to see its status letter: (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if float >= 0 && float < 60 {
					fmt.Println("The status letter for grade ", input, "is F")
				} else if float >= 60 && float < 70 {
					fmt.Println("The status letter for grade ", input, "is D")
				} else if float >= 70 && float < 80 {
					fmt.Println("The status letter for grade ", input, "is C")
				} else if float >= 80 && float < 90 {
					fmt.Println("The status letter for grade ", input, "is B")
				} else if float >= 90 && float <= 100 {
					fmt.Println("The status letter for grade ", input, "is A")
				} else {
					fmt.Println("How exactly did you get that grade???")
				}
			}
		}
	}
}
