package main

import "fmt"

type UserID int

type Num interface {
	~int | ~int8 | ~int16 | ~float32 | ~float64
}

func Add[T Num](a T, b T) T {
	return a + b
}

func main() {
	a := UserID(1)
	b := UserID(2)
	result := Add(a, b)
	fmt.Printf("result: %v\n", result)
}
