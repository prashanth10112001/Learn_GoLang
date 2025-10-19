package main

import (
	"fmt"
	"net/url"
)

const URL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=kgjhmhvbm"
// const URL = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=27"


func main() {
	fmt.Println("Welcome to handling URL")
	fmt.Println(URL)

	//parsing
	result,_ := url.Parse(URL)

	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params is %T\n",qparams)
	fmt.Println(qparams)

	fmt.Println(qparams["coursename"])

	for _,v := range qparams{
		fmt.Println("The values in params is ",v)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath : "user=paisa",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)


}