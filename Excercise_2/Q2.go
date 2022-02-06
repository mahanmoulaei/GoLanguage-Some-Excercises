package main

import (
	"fmt"
)

func main() {
	var inputFloors int
	fmt.Print("Enter the amount of floors of the building : ")
	fmt.Scan(&inputFloors)

	calcBuildHeight(inputFloors)
}

func calcBuildHeight(floors int) {

	var totalHeight int = (4 * floors) + 6
	if floors < 0 {
		fmt.Print("Enter a positive number!\n")
		main()
	} else {
		fmt.Printf("The height of the building is %dm", totalHeight)
	}
}
