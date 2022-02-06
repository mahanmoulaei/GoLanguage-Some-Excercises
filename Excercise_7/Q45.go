package main

import "fmt"

func main() {
	number := fibonacci(10)
	fmt.Println(number)
}

func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
