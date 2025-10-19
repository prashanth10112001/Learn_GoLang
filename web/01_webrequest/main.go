package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"
// const url = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=27"


func main() {
	fmt.Println("LCO web request")
	response,err := http.Get(url)
	checkErrNil(err)

	fmt.Printf("Response i sof type : %T\n",response)
	fmt.Printf("Response i sof type : %v\n",response.StatusCode)

	defer response.Body.Close()

	databytes ,err := ioutil.ReadAll(response.Body)

	checkErrNil(err)

	content := string(databytes)

	fmt.Println(content)
	
}

func checkErrNil(err error)  {
	if err!=nil{
		panic(err)
	}
}