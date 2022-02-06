package main

import "fmt"

func main() {
	for n := 0; n <= 10; n++ {
		fmt.Println("\n***********")
		fmt.Println("Number:", n)
		fmt.Println("Square:", n*n)
		fmt.Println("Cube:", n*n*n)
	}
}
