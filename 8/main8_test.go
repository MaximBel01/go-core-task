package main

import (
	"testing"
	"time"
)

func TestWaitGroupBasic(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(2)

	done := make(chan bool)

	go func() {
		wg.Wait()
		done <- true
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		wg.Done()
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		wg.Done()
	}()

	select {
	case <-done:
	case <-time.After(300 * time.Millisecond):
		t.Error("Wait timed out")
	}
}

func TestWaitGroupNegativeCounter(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for negative counter")
		}
	}()

	wg := NewWaitGroup()
	wg.Add(1)
	wg.Done()
	wg.Done()
}

func TestWaitGroupReuse(t *testing.T) {
	wg := NewWaitGroup()

	wg.Add(2)
	go wg.Done()
	go wg.Done()
	wg.Wait()

	wg.Add(3)
	go wg.Done()
	go wg.Done()
	go wg.Done()
	wg.Wait()
}
