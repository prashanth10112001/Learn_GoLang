package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Welcome to my times")

	// for specific times we need to follow the given format

	presentTime := time.Now()
	fmt.Println("present time: ",presentTime) //2025-10-18 19:58:29.0218255 +0530 IST m=+0.000351901

	fmt.Println(presentTime.Format("01-02-2006")) // 10-18-2025

	fmt.Println(presentTime.Format("01-02-2006 Monday")) // 10-18-2025 Saturday

	fmt.Println(presentTime.Format("01-02-2006 15:04:05  Monday")) // 10-18-2025 20:00:11  Saturday

	createdDate := time.Date(2001,time.November,10,20,23,0,0,time.UTC)
	fmt.Println(createdDate) //2001-11-10 20:23:00 +0000 UTC


	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}