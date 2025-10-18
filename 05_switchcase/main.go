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
}