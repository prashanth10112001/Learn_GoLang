package main

import "fmt"

func main() {
	fmt.Println("Welcometo if loops")

	loginCount := 23

	var result string
	if loginCount <10{
		result = "regular result"
	}else if loginCount>10 {
		result ="something else"
	}else{
		result ="nothing else"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	}else{
		fmt.Println("Number is odd")
	}

	if num := 3; num<10 {
		fmt.Println("Num is less than 10")
	}else {
		fmt.Println("Not less than 10")
	}


}