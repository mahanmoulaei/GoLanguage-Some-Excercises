package main

import (
	"fmt"
	"math"
)

func main() {
	var inputLab float64
	var inputMidterm float64
	var inputFinal float64

	fmt.Print("Enter Laboratory Grade : ")
	fmt.Scan(&inputLab)
	fmt.Print("Enter Midterm Grade : ")
	fmt.Scan(&inputMidterm)
	fmt.Print("Enter Final Exam Grade : ")
	fmt.Scan(&inputFinal)
	calcFinalGrade(inputLab, inputMidterm, inputFinal)

}

func calcFinalGrade(lab float64, midterm float64, final float64) {

	var finalCalc float64 = math.Round((lab*0.40+midterm*0.25+final*0.35)*100) / 100
	if lab > 100 || lab < 0 || midterm > 100 || midterm < 0 || final > 100 || final < 0 {
		fmt.Print("Please enter valid grades!\n")
		main()
	} else {
		fmt.Printf("The final grade is %g", finalCalc)
	}
}
