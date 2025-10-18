package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	greets()
	greet2()

	result := adder(2,3)
	fmt.Println("result is ",result)

	res2,_ := proAdder(1,2,3,4,5)
	fmt.Println("New result is ",res2)

	
}

func adder( a int ,b int) int  {
	return a+b
}

func proAdder( vals ...int) (int ,string) {
	total := 0
	for _,val := range vals{
		total+=val
	}
	return  total,"Hi pro result"
}

func greet2()  {
		fmt.Println("Another greet")
	}


func greets()  {
	fmt.Println("Namaste")
}