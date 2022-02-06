package main

import "fmt"

func main() {

	limit := 15

	i := 1
	result := 1
	for i <= limit {
		if i%2 != 0 {
			result = result * i
		}
		i++
	}

	fmt.Println("the product of all odd numbers between", i, "and", limit, "is:", result)
}
