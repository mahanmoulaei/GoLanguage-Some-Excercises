package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	var sides []float64
	var count int = 0

	var isGood bool

	for {
		if count != 3 {
			fmt.Println("Enter 1 number as side", count+1, "of triangle: (type exit to quit the program)")
			fmt.Scan(&input)
			if input == "exit" {
				os.Exit(0)
			} else {
				if float, error := strconv.ParseFloat(input, 64); error == nil {
					sides = append(sides, float)
					count++
				}
			}
			continue
		} else {
			break
		}

	}
	fmt.Println(sides)

	if sides[0] > sides[1] {
		if sides[0] > sides[2] {
			isGood = check(sides[0], sides[1], sides[2])
		} else { //sides[2] biggest
			isGood = check(sides[2], sides[1], sides[0])
		}
	} else { // sides[1] bigger
		if sides[1] > sides[2] {
			isGood = check(sides[1], sides[0], sides[2])
		} else { //sides[2] biggest
			isGood = check(sides[2], sides[1], sides[0])
		}
	}

	if isGood {
		fmt.Println("The 3 numbers you entered CAN make a triangle!")
	} else {
		fmt.Println("The 3 numbers you entered CANNOT make a triangle!")
	}
}

func check(s1 float64, s2 float64, s3 float64) bool {
	if s1 < (s2 + s3) {
		return true
	} else {
		return false
	}
}
