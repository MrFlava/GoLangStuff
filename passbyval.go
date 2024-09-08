package main

import (
	"fmt"
)
  
func updName(x string) string{
	x = "edge"
	return x
}

func updMenu(y map[string]float64){
	y["coffee" ] = 2.99
}

func main(){
	// group A types -> strings, ints, bools, floats, arrays, structs
	name := "tiffa"

	name = updName(name)

	fmt.Println(name)

	// group B types -> slices, maps, functions
	menu := map[string]float64{
		"pie": 5.95,
		"ice creram": 3.99,
	}

	updMenu(menu)
	fmt.Println(menu)
}