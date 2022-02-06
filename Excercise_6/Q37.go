package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	var number int

	for {
		fmt.Println("Enter a number bigger than 0: (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				if integer > 0 { //check for negatives and 0
					number = integer
					break
				}
			}
		}
	}

	forLoop(number)
	whileLoop(number)

}

func forLoop(number int) {
	isPrime := false
	for i := 1; i < number; i++ {
		if number%i == 0 {
			if i == 1 || i == number {
				isPrime = true
			} else {
				isPrime = false
				break
			}
		}
	}
	if isPrime == true {
		fmt.Println("For Loop =>", number, "is a Prime number.")
	} else {
		fmt.Println("For Loop =>", number, "is NOT a Prime number.")
	}
}

func whileLoop(number int) {
	isPrime := false
	i := 1
	for i < number {
		if number%i == 0 {
			if i == 1 || i == number {
				isPrime = true
			} else {
				isPrime = false
				break
			}
		}
		i++
	}
	if isPrime == true {
		fmt.Println("While Loop =>", number, "is a Prime number.")
	} else {
		fmt.Println("While Loop =>", number, "is NOT a Prime number.")
	}
}
