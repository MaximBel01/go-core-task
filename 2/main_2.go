package main

import (
	"fmt"
	"math/rand"
)

func generateRandomSlice(n int) []int {
	slice := make([]int, n)
	for i := range n {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func sliceExample(s []int) []int {
	even := make([]int, 0)
	for _, num := range s {
		if num%2 == 0 {
			even = append(even, num)
		}
	}
	return even
}

func addElements(s []int, num int) []int {
	return append(s, num)
}

func copySlice(s []int) []int {
	newSlice := make([]int, len(s))
	copy(newSlice, s)
	return newSlice
}

func removeElement(s []int, index int) ([]int, error) {
	if index < 0 || index >= len(s) {
		return nil, fmt.Errorf("index out of range")
	}
	newSlice := make([]int, 0, len(s)-1)
	newSlice = append(newSlice, s[:index]...)
	newSlice = append(newSlice, s[index+1:]...)
	return newSlice, nil
}

func main() {
	n := 10   //how many generate
	add := 99 //number to add
	rm := 3   //index to remove

	originalSlice := generateRandomSlice(n)
	fmt.Printf("Original Slice: \t%v\n", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Printf("Even Numbers: \t\t%v\n", evenSlice)

	addedSlice := addElements(originalSlice, add)
	fmt.Printf("After Adding %v: \t%v\n", add, addedSlice)

	copiedSlice := copySlice(originalSlice)
	originalSlice[0] = -1
	fmt.Printf("Original Modified: \t%v\n", originalSlice)
	fmt.Printf("Copied Slice: \t\t%v\n", copiedSlice)

	removedSlice, err := removeElement(originalSlice, rm)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("After Removing Index %v: %v\n", rm, removedSlice)
	}
}
