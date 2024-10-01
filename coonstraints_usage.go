package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func getBiggerNumber[T numbers](t1, t2 T) T {
	if t1 > t2 {
		return t1
	}
	return t2
}

func getBiggerNumberWithComparable[T constraints.Ordered](t1, t2 T) T {
	if t1 > t2 {
		return t1
	}
	return t2
}

func main() {
	fmt.Println(getBiggerNumber(2.5, -4.0))
	fmt.Println(getBiggerNumberWithComparable(2.5, -4.0))
}
