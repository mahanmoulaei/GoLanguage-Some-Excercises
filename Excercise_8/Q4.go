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

	for {
		fmt.Println("Enter the third number : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				num3 = integer
				break
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	fmt.Println("The Sum of", num1, "&", num2, "&", num3, "is equal to", Sum(num1, num2, num3))
	fmt.Println("The Avarage of", num1, "&", num2, "&", num3, "is equal to ", Average(num1, num2, num3))
	fmt.Println("The Product of", num1, "&", num2, "&", num3, "is equal to ", Product(num1, num2, num3))
	fmt.Println("The Smallest number between", num1, "&", num2, "&", num3, "is ", getSmallestNumber(num1, num2, num3))
	fmt.Println("The Largest number between", num1, "&", num2, "&", num3, "is ", getLargestNumber(num1, num2, num3))

}

func Sum(num1 int, num2 int, num3 int) int {
	return num1 + num2 + num3
}

func Average(num1 int, num2 int, num3 int) int {
	return (num1 + num2 + num3) / 3
}

func Product(num1 int, num2 int, num3 int) int {
	return num1 * num2 * num3
}

func getLargestNumber(num1 int, num2 int, num3 int) int {
	if num1 > num2 && num1 > num3 {
		return num1
	} else if num2 > num1 && num2 > num3 {
		return num2
	} else {
		return num3
	}
}

func getSmallestNumber(num1 int, num2 int, num3 int) int {
	if num1 < num2 && num1 < num3 {
		return num1
	} else if num2 < num1 && num2 < num3 {
		return num2
	} else {
		return num3
	}
}
