package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	var num, digit1, digit2, digit4, digit5 int

	for {
		fmt.Println("Enter a 5 digit number : (type exit to quit the program):")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				if integer >= 10000 && integer <= 99999 {
					num = integer
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	digit1 = num / 10000
	digit2 = num / 1000 % 10
	digit4 = num / 10 % 10
	digit5 = num % 10

	if digit1 == digit5 && digit2 == digit4 {
		fmt.Println("Your number is a Palindrome..")
	} else if digit1 != digit5 || digit2 != digit4 {
		fmt.Println("Your number is Not a palindrome..")
	}

}
