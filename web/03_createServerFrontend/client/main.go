package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)



func main() {
	fmt.Println("Welcome to web verb")
	// performGetReq()
	// performJsonPostReq()
	performPostFormReq()
}

func performGetReq(){
	const myurl = "http://localhost:3000/get"

	res,err := http.Get(myurl)

	checkNilErr(err)

	defer res.Body.Close()

	fmt.Println("Status code: ",res.StatusCode)
	fmt.Println("Content length: ",res.ContentLength)

	var resString strings.Builder
	content , _ := ioutil.ReadAll(res.Body)
	byteCount , _ := resString.Write(content)

	fmt.Println("Byte count is ",byteCount)
	fmt.Println(resString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))

}

func performJsonPostReq(){
	const myurl = "http://localhost:3000/post"

	//fake json payload

	reqBody := strings.NewReader(`
		{
			"coursename":"lets's go with go lang",
			"price":2000,
			"platform":"learncode.dev"
		}
	`)

	res,err := http.Post(myurl,"application/json",reqBody)

	checkNilErr(err)

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	checkNilErr(err)

	fmt.Println(string(content))
}

func performPostFormReq()  {
	const myurl ="http://localhost:3000/postform"

	data := url.Values{}

	data.Add("firstname","paisa")
	data.Add("age","20")
	data.Add("roll","33")
	data.Add("email","as@outlook.in")

	res,err := http.PostForm(myurl,data)
	checkNilErr(err)

	defer res.Body.Close()
	content,err := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))
	
}

func checkNilErr(err error){
	if err!=nil{
		panic(err)
	}
}