package main

import (
	"fmt"
)

func updNameHere(x *string){
	*x = "edge"
}

func main(){
	name := "tifa"

	// updNameHere(name)

	// fmt.Println("memory address: ", &name)


	m := &name
	// fmt.Println("memory address: ", m)w
	// fmt.Println("val at mem: ", *m)

	fmt.Println(name)
	updNameHere(m)
	fmt.Println(name)

	// fmt.Println(name)
}