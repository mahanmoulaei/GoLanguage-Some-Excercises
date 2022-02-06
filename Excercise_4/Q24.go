package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Print("You tossed 2 coins!\n")
	coin1 := []string{
		"heads",
		"tails",
	}

	coin2 := []string{
		"heads",
		"tails",
	}

	rand.Seed(time.Now().UnixNano())

	side1 := coin1[rand.Intn(len(coin1))]
	side2 := coin2[rand.Intn(len(coin1))]

	if side1 == "heads" && side2 == "heads" {
		fmt.Print("Both coins landed on ", side1, ", you win $10!")
	} else if side1 == "heads" && side2 == "tails" {
		fmt.Print("First coin landed on ", side1, " second coin landed on ", side2, ", you win $5!")
	} else {
		fmt.Print("You lose :(")
	}
}
