package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to int conversion")
	fmt.Println("Please arte between 1 to 5:-")

	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ",input)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil{
		fmt.Println("err:- ",err)

	}else{

		fmt.Println("Added one to your rating:- ",numrating +1)

	}

	a := 'a'
	fmt.Println("a is ",a)

}