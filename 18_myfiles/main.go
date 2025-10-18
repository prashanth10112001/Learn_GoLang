package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to   files")

	content := "this needs to go int a file - learning .com.in"

	file,err := os.Create("./mylcofile.txt")

	checkNilError(err)

	length,err := io.WriteString(file,content)

	checkNilError(err)

	fmt.Println("Length is ,",length)
	defer file.Close()
	readFile("./mylcofile.txt")
}

func readFile(filename string){
	databyte ,err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file is \n",string(databyte))
}

func checkNilError(err error){
	if err != nil{
		panic(err)
	}
}