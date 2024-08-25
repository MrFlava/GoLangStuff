package main

import "fmt"

func main(){
	var ages [3]int = [3]int{20, 25, 30}

	names := [4]string{"youshi", "mario", "peach", "bowser"}
	names[1] = "lou"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	var scores = []int{100, 500, 60}
	scores[2] = 25

	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)
}