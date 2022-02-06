package main

import (
	"fmt"
)

func main() {

	var input1 float64
	var input2 float64

	fmt.Print("Enter value 1 : ")
	fmt.Scan(&input1)
	fmt.Print("Enter value 2 : ")
	fmt.Scan(&input2)

	greaterVal(input1, input2)
}

func greaterVal(num1 float64, num2 float64) {

	if num1 > num2 {
		fmt.Printf("The greater value is : %g", num1)
	} else {
		fmt.Printf("The greater value is : %g", num2)
	}

}
