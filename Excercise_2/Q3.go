package main

import (
	"fmt"
	"math"
)

func main() {

	var inputKpa float64
	fmt.Print("Enter the amount in kilopascal units(KPa) : ")
	fmt.Scan(&inputKpa)

	convertToAtm(inputKpa)
}

func convertToAtm(kpa float64) {
	var atm float64 = math.Round(kpa/101.325*10000) / 10000
	fmt.Printf("%g kPa(kilopascal) is equal to %g Atm(atmospheric pressure)", kpa, atm)
}
