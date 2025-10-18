package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers chapter")

	var ptr *int
	fmt.Println("Value of pointer is, ",ptr)

	myNum := 23
	var ptrToNum = &myNum
	fmt.Println("Value of actual ptr address is, ",ptrToNum)
	fmt.Println("Value of actual ptr is, ",*ptrToNum)

	*ptrToNum = *ptrToNum * 2
	fmt.Println("New value is, ",*ptrToNum)
	fmt.Println("New value is, ",myNum)

}