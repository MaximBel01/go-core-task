package main

import (
	"fmt"
	"sort"
)

func FindIntersection(a, b []int) (bool, []int) {
	elements := make(map[int]bool)
	for _, num := range a {
		elements[num] = true
	}

	intersection := []int{}
	for _, num := range b {
		if elements[num] {
			intersection = append(intersection, num)
			elements[num] = false
		}
	}
	return len(intersection) > 0, intersection
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasIntersection, intersection := FindIntersection(a, b)

	sort.Ints(intersection)

	fmt.Printf("Есть ли пересечения: %t\n", hasIntersection)
	fmt.Printf("Общие элементы: %v\n", intersection)
}
