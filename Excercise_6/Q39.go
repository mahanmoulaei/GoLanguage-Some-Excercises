package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//forLoop()
	whileLoop()
}

func forLoop() {
	var input string
	var array []float64
	sum := 0.0
	limit := 1000
	var firstTime bool = true
	for i := 0; i < limit; i++ {
		if firstTime {
			fmt.Println("Position", i+1, "- Enter a grade : (type exit to quit the program)")
		} else {
			fmt.Println("Position", i+1, "- Enter another grade : (type exit to quit the program) (to quit the program and view the average until now, type \"-1-1\" WITHOUT the quotations)")
		}
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else if input == "-1-1" {
			break
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if finalFloat, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float), 64); err == nil {
					array = append(array, finalFloat)
					sum += finalFloat
					if firstTime {
						firstTime = false
					}
					continue
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	fmt.Println(array)
	fmt.Println("Average is", (sum / float64(len(array))))

}

func whileLoop() {
	var input string
	var array []float64
	sum := 0.0
	i := 0
	limit := 1000
	var firstTime bool = true
	for i < limit {
		if firstTime {
			fmt.Println("Position", i+1, "- Enter a grade : (type exit to quit the program)")
		} else {
			fmt.Println("Position", i+1, "- Enter another grade : (type exit to quit the program) (to quit the program and view the average until now, type \"-1-1\" WITHOUT the quotations)")
		}
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else if input == "-1-1" {
			break
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if finalFloat, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float), 64); err == nil {
					array = append(array, finalFloat)
					sum += finalFloat
					if firstTime {
						firstTime = false
					}
					i++
					continue
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	fmt.Println(array)
	fmt.Println("Average is", (sum / float64(len(array))))

}
