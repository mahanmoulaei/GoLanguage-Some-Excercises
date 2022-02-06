package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	var scores []int
	var count int = 0

	for {
		if count != 6 {
			//fmt.Println("Enter score No.", count+1, " (type exit to quit the program)")
			fmt.Printf("Enter score No.%d (type exit to quit the program)\n", count+1)
			fmt.Scan(&input)
			if input == "exit" {
				os.Exit(0)
			} else {
				if integer, error := strconv.Atoi(input); error == nil {
					if integer >= 0 { //check for negatives
						scores = append(scores, integer)
						count++
					}
				}
			}
			continue
		} else {
			break
		}
	}

	min, max, sum := findMinAndMaxAndSumOf(scores)

	fmt.Printf("The final score is %d. (Min score of %d and Max score of %d have been eliminated)", ((sum - min - max) / 4), min, max)
}

func findMinAndMaxAndSumOf(array []int) (min int, max int, sum int) {
	min = array[0]
	max = array[0]
	for _, value := range array {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
		sum += value
	}
	return min, max, sum
}
