package main

import (
	"fmt"
)

func main(){
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is:", i)
	// }

	names := []string{"mario", "luigi", "aaaa"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Println(index, value)
	// }

	for _, value := range names {
		fmt.Printf("the val is %v \n", value)
		value = "new"
	}

	fmt.Println(names)
 
}