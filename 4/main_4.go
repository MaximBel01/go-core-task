package main

import (
	"fmt"
)

func Difference(slice1, slice2 []string) []string {
	elementsInSlice2 := make(map[string]bool)
	for _, item := range slice2 {
		elementsInSlice2[item] = true
	}

	result := []string{}

	for _, item := range slice1 {
		if !elementsInSlice2[item] {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	diff := Difference(slice1, slice2)

	fmt.Println("Elements in slice1 but not in slice2:", diff)
}
