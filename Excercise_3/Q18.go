package main

import (
	"fmt"
)

func main() {

	var num1 int
	var num2 int

	fmt.Print("Check if number 1 is a multiple of number 2\n")
	fmt.Print("Enter number 1 : ")
	fmt.Scan(&num1)
	fmt.Print("Enter number 2 : ")
	fmt.Scan(&num2)

	if num1%num2 == 0 {
		fmt.Println(num1, "is a multiple of", num2)
	} else {
		fmt.Println(num1, "is not a multiple of", num2)
	}
}
