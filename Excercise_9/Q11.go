package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	var num1, num2, num3 int

	for {
		fmt.Println("Enter the first non zero integer number : (type exit to quit the program):")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				if integer > 0 {
					num1 = integer
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	for {
		fmt.Println("Enter the second non zero integer number : (type exit to quit the program):")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				if integer > 0 {
					num2 = integer
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	for {
		fmt.Println("Enter the third non zero integer number : (type exit to quit the program):")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				if integer > 0 {
					num3 = integer
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	if canFormRightAngledTriangle(num1, num2, num3) {
		fmt.Println("Your numbers CAN form the sides of a right-angled triangle...")
	} else {
		fmt.Println("Your numbers CANNOT form the sides of a right-angled triangle...")
	}
}

func canFormRightAngledTriangle(num1 int, num2 int, num3 int) bool {
	var answer bool
	if (num1*num1) == (num2*num2)+(num3*num3) || (num2*num2) == (num1*num1)+(num3*num3) || (num3*num3) == (num1*num1)+(num2*num2) {
		answer = true
	} else {
		answer = false
	}
	return answer
}
