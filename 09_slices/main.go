package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruits = []string{}
	fmt.Printf("Type of fruits is %T , ",fruits)

	var veggies = []string{"apple","veggie"}
	fmt.Printf("Type of veggies is %T \n",veggies)

	veggies= append(veggies, "maggie","yippie")
	fmt.Println(veggies)

	veggies = append(veggies[1:3])
	fmt.Println(veggies)

	highScores := make([]int, 4)

	highScores[0] =234
	highScores[1] =189
	highScores[2] =137
	highScores[3] =200

	fmt.Println(highScores)

	lowScores := make([]int, 4)

	lowScores[0] =234
	lowScores[1] =189
	lowScores[2] =137
	lowScores[3] =200

	lowScores = append(lowScores, 678,987,980)

	fmt.Println(lowScores)

	sort.Ints(lowScores)
	fmt.Println(lowScores)
	fmt.Println(sort.IntsAreSorted(lowScores))

	var courses = []string{"js","react","html","r","c"}
	fmt.Println(courses)

	var index int =2
	courses = append(courses[:index],courses[index+1:]... )
	fmt.Println(courses)

}