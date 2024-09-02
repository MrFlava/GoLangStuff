package main

import (
	"fmt"
	"strings"
	"sort"
)

func main(){
	greeting := "hello!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ToUpper(greeting))

	ages := []int{45, 12, 234, 46, 34, 8}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 34)
	fmt.Println(index)

	names := []string{"cool", "frank", "sofia", "apple"}
	sort.Strings(names)
	fmt.Println(names)
}