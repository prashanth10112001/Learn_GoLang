package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in go lang")
	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6)+1
	fmt.Println("Value of dice is:- ",diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
		
	case 5:
		fmt.Println("Five")
		fallthrough
	case 6:
		fmt.Println("Six")
	default:
		fmt.Println("default")
	}


}