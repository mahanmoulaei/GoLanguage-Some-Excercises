package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	var num int

	for {
		fmt.Println("Enter a positive integer number : (type exit to quit the program):")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				if integer > 0 {
					num = integer
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	f := factorial(num)

	fmt.Println("Factorial of " + strconv.Itoa(num) + " is: " + strconv.Itoa(f))
}

func factorial(number int) int {
	result := 1

	for i := 2; i <= number; i++ {
		result = result * i
	}
	return result
}
