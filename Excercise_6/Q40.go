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
	var array []int
	var min int
	var max int
	limit := 1000
	var firstTime bool = true
	for i := 0; i < limit; i++ {
		if firstTime {
			fmt.Println("Position", i+1, "- Enter a number : (type exit to quit the program)")
		} else {
			fmt.Println("Position", i+1, "- Enter another number : (type exit to quit the program) (to quit the program and view the Min & Max until now, type \"-1-1\" WITHOUT the quotations)")
		}
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else if input == "-1-1" {
			break
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				array = append(array, integer)

				if firstTime {
					min = integer
					max = integer
					firstTime = false
				} else {
					if min > integer {
						min = integer
					}
					if max < integer {
						max = integer
					}
				}
				continue
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	fmt.Println(array)
	fmt.Println("Min is", min, "& Max is", max)

}

func whileLoop() {
	var input string
	var array []int
	var min int
	var max int
	i := 0
	limit := 1000
	var firstTime bool = true
	for i < limit {
		if firstTime {
			fmt.Println("Position", i+1, "- Enter a number : (type exit to quit the program)")
		} else {
			fmt.Println("Position", i+1, "- Enter another number : (type exit to quit the program) (to quit the program and view the Min & Max until now, type \"-1-1\" WITHOUT the quotations)")
		}
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else if input == "-1-1" {
			break
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				array = append(array, integer)
				if firstTime {
					min = integer
					max = integer
					firstTime = false
				} else {
					if min > integer {
						min = integer
					}
					if max < integer {
						max = integer
					}
				}
				i++
				continue
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	fmt.Println(array)
	fmt.Println("Min is", min, "& Max is", max)

}
