package main

import (
	"fmt"
)

func main() {

	var inputnum int
	fmt.Print("Enter a number : ")
	fmt.Scan(&inputnum)

	evenOrOdd(inputnum)
}

func evenOrOdd(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is an odd number")
	}
}
