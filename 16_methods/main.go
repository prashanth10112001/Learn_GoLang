package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods")

	paisa := User{"paisa","p@.com",true,21}
	fmt.Printf("My values are %+v\n",paisa)

	paisa.GetStatus()
	paisa.NewMail()
	fmt.Printf("My email value are %v\n",paisa.email)

}

type User struct{
	Name string
	email string
	Status bool
	Age int

}

func (u User) GetStatus()  {
	fmt.Println("Is user active: ",u.Status)
}

func (u User) NewMail()  {
	u.email = "asd@gmail.com"
	fmt.Println("Email of user is : ",u.email)
}