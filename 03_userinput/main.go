package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the reading for a pizza:-")

	// comma ok || err ok -> syntax
	input , _ := reader.ReadString('\n')
	fmt.Println("thanks for input, ",input)
	fmt.Printf("Type for input is %T ",input)


}