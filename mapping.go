package main

import (
	"fmt"
)

type Num interface {
	int | int8 | int16 | float32 | float64
}

func MapValues[T Num](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}

	return newValues
}

func main() {
	result := MapValues([]float64{1.1, 2.2, 3.3}, func(n float64) float64 {
		return n * 2
	})
	fmt.Printf("result: %+v\n", result)
}
