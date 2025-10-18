package main

import "fmt"

// outside not allowed
// jwtToken := 30000
// allowed
// var s int = 4

// This is public variable by Captial letter
const LoginToken string = "gjbshuiak"

func main()  {
	fmt.Print("Variables")

	var username string = "Paisa"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isverified bool = true
	fmt.Println(isverified)
	fmt.Printf("Variable is of type: %T \n", isverified)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	var vals int = 255
	fmt.Println(vals)
	fmt.Printf("Variable is of type: %T \n", vals)

	var smallFloat float32 = 255.22232334223232344232
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var bigFloat float64 = 255.22232334223232344232
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)

	var anotherVal int
	fmt.Println(anotherVal)
	fmt.Printf("Variable is of type: %T \n", anotherVal)

	var website = "helloworld.com"
	fmt.Println(website)

	// website = 3 // error once string always a string

	numberOfuser := 30000
	fmt.Println(numberOfuser)


	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}