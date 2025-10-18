package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array")

	var fruits [4]string

	fruits[0]= "mango"
	fruits[1]= "banana"
	fruits[3]= "peech"

	fmt.Println("Fruit list is : -", fruits)
	fmt.Println("Fruit list length is : -", len(fruits))

	var vegies = [3]string{"potato","mushroom"} 
	fmt.Println("veggies list is ",len(vegies))
}