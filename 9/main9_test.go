package main

import (
	"math"
	"testing"
)

func TestCubeConverter(t *testing.T) {
	t.Run("Basic conversion", func(t *testing.T) {
		input := make(chan uint8)
		output := CubeConverter(input)

		go func() {
			defer close(input)
			input <- 2
			input <- 3
		}()

		expected := []float64{8, 27}
		for _, exp := range expected {
			res := <-output
			if res != exp {
				t.Errorf("Expected %.2f, got %.2f", exp, res)
			}
		}

		if _, ok := <-output; ok {
			t.Error("Output channel should be closed")
		}
	})

	t.Run("Empty input", func(t *testing.T) {
		input := make(chan uint8)
		close(input)

		output := CubeConverter(input)
		if _, ok := <-output; ok {
			t.Error("Output channel should be closed")
		}
	})

	t.Run("Max uint8 value", func(t *testing.T) {
		input := make(chan uint8)
		output := CubeConverter(input)

		go func() {
			defer close(input)
			input <- 255
		}()

		expected := math.Pow(255, 3)
		res := <-output
		if res != expected {
			t.Errorf("Expected %.2f, got %.2f", expected, res)
		}
	})
}
