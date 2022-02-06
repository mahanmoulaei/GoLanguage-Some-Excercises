package main

import (
	"fmt"
)

type Sequence struct {
	min       int
	max       int
	increment int
}

func main() {
	A := Sequence{min: 5, max: 40, increment: 5}
	B := Sequence{min: 3, max: 15, increment: 2}
	C := Sequence{min: 80, max: 20, increment: -10}
	D := Sequence{min: 1, max: 720, increment: 1}
	fmt.Println("\n\t*** A ***")
	letterAForLoop(A)
	letterAWhileLoop(A)
	fmt.Println("\n\t*** B ***")
	letterBForLoop(B)
	letterBWhileLoop(B)
	fmt.Println("\n\t*** C ***")
	letterCForLoop(C)
	letterCWhileLoop(C)
	fmt.Println("\n\t*** D ***")
	letterDForLoop(D)
	letterDWhileLoop(D)
}

func letterAForLoop(algorithm Sequence) {
	var array []int
	fmt.Println("*** Using For Loop ***")
	for i := algorithm.min; i <= algorithm.max; i += algorithm.increment {
		array = append(array, int(i))
	}
	fmt.Println(array)
}

func letterBForLoop(algorithm Sequence) {
	var array []int
	fmt.Println("*** Using For Loop ***")
	for i := algorithm.min; i <= algorithm.max; i += algorithm.increment {
		array = append(array, int(i))
	}
	fmt.Println(array)
}

func letterCForLoop(algorithm Sequence) {
	var array []int
	fmt.Println("*** Using For Loop ***")
	for i := algorithm.min; i >= algorithm.max; i += algorithm.increment {
		array = append(array, int(i))
	}
	fmt.Println(array)
}

func letterDForLoop(algorithm Sequence) {
	var array []int
	fmt.Println("*** Using For Loop ***")
	for i := algorithm.min; i <= algorithm.max; i *= algorithm.increment {
		array = append(array, int(i))
		algorithm.increment++
	}
	fmt.Println(array)
}

func letterAWhileLoop(algorithm Sequence) {
	var array []int

	fmt.Println("*** Using While loop ***")
	for algorithm.min <= algorithm.max {
		array = append(array, int(algorithm.min))
		algorithm.min += algorithm.increment
	}
	fmt.Println(array)
}

func letterBWhileLoop(algorithm Sequence) {
	var array []int

	fmt.Println("*** Using While loop ***")
	for algorithm.min <= algorithm.max {
		array = append(array, int(algorithm.min))
		algorithm.min += algorithm.increment
	}
	fmt.Println(array)
}

func letterCWhileLoop(algorithm Sequence) {
	var array []int

	fmt.Println("*** Using While loop ***")
	for algorithm.min >= algorithm.max {
		array = append(array, int(algorithm.min))
		algorithm.min += algorithm.increment
	}
	fmt.Println(array)
}

func letterDWhileLoop(algorithm Sequence) {
	var array []int

	fmt.Println("*** Using While loop ***")
	for algorithm.min <= algorithm.max {
		array = append(array, int(algorithm.min))
		algorithm.min *= algorithm.increment
		algorithm.increment++
	}
	fmt.Println(array)
}
