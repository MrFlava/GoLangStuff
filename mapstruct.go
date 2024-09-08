package main

import (
	"fmt"
)

func main(){
	menu := map[string]float64{
		"soup": 1.337,
		"pie": 7.99,
		"taylor swift cake": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])
	
	// looping maps
	for k, v := range menu {
		fmt.Println(k,"-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		214124132: "sonya",
		324121248: "vlad",
		184013413: "fmtprintln", 
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[214124132])

	phonebook[184013413] = "bowser"
	fmt.Println(phonebook)

	phonebook[1] = "coolest"
	fmt.Println(phonebook)


}
