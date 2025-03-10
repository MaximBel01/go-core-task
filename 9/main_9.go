package main

import (
	"fmt"
	"math"
)

func CubeConverter(input <-chan uint8) <-chan float64 {
	output := make(chan float64)
	go func() {
		defer close(output)
		for num := range input {
			output <- math.Pow(float64(num), 3)
		}
	}()
	return output
}

func main() {
	input := make(chan uint8)
	go func() {
		defer close(input)
		numbers := []uint8{2, 3, 5, 10}
		for _, num := range numbers {
			input <- num
		}
	}()

	output := CubeConverter(input)
	for result := range output {
		fmt.Printf("%.2f\n", result)
	}
}
