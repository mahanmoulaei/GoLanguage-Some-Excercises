package main

import (
	"fmt"
	"math"
)

func main() {

	var inputHeight float64
	var inputWidth float64
	var inputLength float64

	fmt.Print("Enter rectangular prism height : ")
	fmt.Scan(&inputHeight)
	fmt.Print("Enter rectangular prism width : ")
	fmt.Scan(&inputWidth)
	fmt.Print("Enter rectangular prism length : ")
	fmt.Scan(&inputLength)

	CalculateVolume(inputHeight, inputWidth, inputLength)
}

func CalculateVolume(height float64, width float64, length float64) {

	var volume float64 = math.Round(length*width*height*100) / 100

	if height <= 0 || width <= 0 || length <= 0 {
		fmt.Print("Please enter valid measurements!\n")
		main()
	} else {
		fmt.Printf("The volume of your rectangular prism is %g", volume)
	}
}
