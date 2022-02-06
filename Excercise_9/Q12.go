///I DON'T KNOW HOW TO DECRYPT!!
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	var originalNumber int

	for {
		fmt.Println("Enter a 4 digit number to get it encrypted : (type exit to quit the program):")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				if integer >= 1000 && integer <= 9999 {
					originalNumber = integer
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	encrypted := encrypt(originalNumber)
	fmt.Println("Encrypted number is", encrypted)
	/*
		for {
			fmt.Println("Enter a 4 digit number to get it decrypted : (type exit to quit the program):")
			fmt.Scan(&input)
			if input == "exit" {
				os.Exit(0)
			} else {
				if integer, error := strconv.Atoi(input); error == nil {
					if integer >= 1000 && integer <= 9999 {
						originalNumber = integer
						break
					}
				}
				fmt.Println("Error In The Entered Number! Try Again...")
			}
		}

		decrypted := encrypt(originalNumber)
		fmt.Println("Decrypted number is", decrypted)
	*/
}

func encrypt(value int) int {
	var digit1, digit2, digit3, digit4 int

	digit1 = value / 1000
	digit2 = (value % 1000) / 100
	digit3 = (value % 100) / 10
	digit4 = value % 10
	fmt.Println(digit1, digit2, digit3, digit4)

	digit1 = (digit1 + 7) % 10
	digit2 = (digit2 + 7) % 10
	digit3 = (digit3 + 7) % 10
	digit4 = (digit4 + 7) % 10

	encryptedNumber := (digit3 * 1000) + (digit4 * 100) + (digit1 * 10) + (digit2)
	return encryptedNumber
}

/*
9 - 6
8 - 5
7 - 4
6 - 3
5 - 2
4 - 1
3 - 0
2 - 2
1 - 1
*/
/*
func decrypt(value int) int {
	var digit1, digit2, digit3, digit4 int
	var array []int
	digit1 = value / 1000
	digit2 = (value % 1000) / 100
	digit3 = (value % 100) / 10
	digit4 = value % 10

	fmt.Println(digit1, digit2, digit3, digit4)

	array = append(array, digit1)
	array = append(array, digit2)
	array = append(array, digit3)
	array = append(array, digit4)

	for i := 1; i <= len(array); i++ {
		switch array[i] {
		case 0:
			if i == 1 {
				digit1 = 3
			} else if i == 2 {
				digit2 = 3
			} else if i == 3 {
				digit3 = 3
			} else if i == 4 {
				digit4 = 3
			}
		}
		case 1:
			if i == 1 {
				digit1 = 1
			} else if i == 2 {
				digit2 = 1
			} else if i == 3 {
				digit3 = 1
			} else if i == 4 {
				digit4 = 1
			}
		}
		case 2:
			if i == 1 {
				digit1 = 5
			} else if i == 2 {
				digit2 = 5
			} else if i == 3 {
				digit3 = 5
			} else if i == 4 {
				digit4 = 5
			}
		}
		case 3:
			if i == 1 {
				digit1 = 6
			} else if i == 2 {
				digit2 = 6
			} else if i == 3 {
				digit3 = 6
			} else if i == 4 {
				digit4 = 6
			}
		}
		case 4:
			if i == 1 {
				digit1 = 7
			} else if i == 2 {
				digit2 = 7
			} else if i == 3 {
				digit3 = 7
			} else if i == 4 {
				digit4 = 7
			}
		}
		case 5:
			if i == 1 {
				digit1 = 8
			} else if i == 2 {
				digit2 = 8
			} else if i == 3 {
				digit3 = 8
			} else if i == 4 {
				digit4 = 8
			}
		}
		case 6:
			if i == 1 {
				digit1 = 9
			} else if i == 2 {
				digit2 = 9
			} else if i == 3 {
				digit3 = 9
			} else if i == 4 {
				digit4 = 9
			}
		}
		case 7:
			if i == 1 {
				digit1 = 3
			} else if i == 2 {
				digit2 = 3
			} else if i == 3 {
				digit3 = 3
			} else if i == 4 {
				digit4 = 3
			}
		}
		case 8:

		}
		case 9:

		}
	}

	decryptedNumber := (digit3 * 1000) + (digit4 * 100) + (digit1 * 10) + (digit2)
	return decryptedNumber
}
*/
