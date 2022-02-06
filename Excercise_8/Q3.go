package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	var num1, num2 int

	for {
		fmt.Println("Enter the first number : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				num1 = integer
				break
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	for {
		fmt.Println("Enter the second number : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				num2 = integer
				break
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	if isEqual(num1, num2) {
		fmt.Println(num1, "is equal to", num2)
	} else {
		min, max := getMinMax(num1, num2)
		fmt.Println(max, "is greater than", min)
	}

}

func getMinMax(num1 int, num2 int) (int, int) {
	var min, max int
	if num1 < num2 {
		min = num1
		max = num2
	} else {
		min = num2
		max = num1
	}
	return min, max
}

func isEqual(num1 int, num2 int) bool {
	if num1 == num2 {
		return true
	} else {
		return false
	}

}
