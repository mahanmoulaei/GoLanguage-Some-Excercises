package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	var limitMin int
	var limitMax int

	for {
		fmt.Println("Enter the Min Number(inclusive) : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				limitMin = integer
				break
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	for {
		fmt.Println("Enter the Max Number(inclusive) : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				limitMax = integer
				break
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	fmt.Printf("For Loop => The sum of the integers between %d and %d, inclusively, is %d\n", limitMin, limitMax, forLoop(limitMin, limitMax))
	fmt.Printf("While Loop => The sum of the integer between %d and %d, inclusively, is %v\n", limitMin, limitMax, whileLoop(limitMin, limitMax))

}

func forLoop(limitMin int, limitMax int) int {
	var sum int
	for i := limitMin; i <= limitMax; i++ {
		sum += i
	}
	return sum
}

func whileLoop(limitMin int, limitMax int) int {
	var sum int
	for limitMin <= limitMax {
		sum += limitMin
		limitMin++
	}
	return sum
}
