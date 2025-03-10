package main

import (
	"testing"
	"time"
)

func TestMergeChannels(t *testing.T) {
	t.Run("Basic merge", func(t *testing.T) {
		ch1 := make(chan int)
		ch2 := make(chan int)
		ch3 := make(chan int)

		merged := MergeChannels(ch1, ch2, ch3)

		go func() {
			ch1 <- 1
			ch2 <- 2
			ch3 <- 3
			close(ch1)
			close(ch2)
			close(ch3)
		}()

		expected := map[int]bool{1: true, 2: true, 3: true}
		result := make(map[int]bool)

		for num := range merged {
			result[num] = true
		}

		if !mapsEqual(expected, result) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("With closed channels", func(t *testing.T) {
		ch1 := make(chan int)
		ch2 := make(chan int)
		close(ch2)
		merged := MergeChannels(ch1, ch2)

		go func() {
			ch1 <- 42
			close(ch1)
		}()

		select {
		case num := <-merged:
			if num != 42 {
				t.Errorf("Expected 42, got %d", num)
			}
		case <-time.After(1 * time.Second):
			t.Error("Timeout: no data received")
		}

		_, ok := <-merged
		if ok {
			t.Error("Channel should be closed")
		}
	})
}

func mapsEqual(a, b map[int]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if !b[k] {
			return false
		}
	}
	return true
}
