package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	days := []string{"Sunday", "Tuesday", "Wednesday"}

	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days{
		fmt.Printf("Index is %v and val is %v\n",index,day)
	}

	for _, day := range days{
		fmt.Printf("Index  and val is %v\n",day)
	}

	rough := 1

	for rough<10{
		if rough==3{
			goto lco
		}else if rough ==7{
			break
		}else if rough ==5 {
			rough++
			continue
		}
		fmt.Println("value is ", rough)
		rough++
	}

	lco:
		fmt.Println("Jumping at learn")
}