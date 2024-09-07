package main

import (
	"fmt"
)

func main(){
	age:=25

	fmt.Println(age <=50)
	fmt.Println(age >=50)
	fmt.Println(age ==45)
	fmt.Println(age !=50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")		
	} else {
		fmt.Println("age is less than 45")		
	}

	names := []string{"mar", "io", "foo", "bar", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("dali at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking news at pos", index)
			break
		}

		fmt.Printf("THE VAlue as pos %v is %v \n", index, value)
	}
}