package main

import "fmt"

func main(){
	age := 35
	name := "shaun"

	fmt.Print("hello, ")
	fmt.Print("world \n")
	fmt.Print("new line \n")

	fmt.Println("hello ninjas")
	fmt.Println("goodbye ninjas")
	fmt.Println("my age is", age)

	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("score %f \n", 2.53442)

	var cool = fmt.Sprintf("my coolest %v \n", age)
	fmt.Println("saved:", cool)
}