package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("hello world - 1")
	defer fmt.Println("hello world - 2")
	fmt.Println("Welcmoe to defers")
	defer mydefer()
	defer fmt.Println("hello world - 3")
}

func  mydefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i," ")
	}
}