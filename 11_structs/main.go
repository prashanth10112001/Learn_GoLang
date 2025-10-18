package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	paisa := User{"Paisa","paisa@gmail.com",21,true}
	fmt.Println(paisa)
	fmt.Printf("Paisa details with age %v are %+v\n",paisa.Age,paisa)
}

// no inheritance ; no super ; no parent ; no child ;
type User struct {
	Name   string
	email  string
	Age    int
	Status bool
}