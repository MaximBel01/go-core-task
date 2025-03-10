package main

import (
	"context"
	"testing"
	"time"
)

func TestRandomGenerator(t *testing.T) {
	t.Run("Basic functionality", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch := RandomGenerator(ctx)

		for range 5 {
			select {
			case num := <-ch:
				if num < 0 || num >= 100 {
					t.Errorf("Number %d is out of range [0, 100)", num)
				}
			case <-time.After(1 * time.Second):
				t.Fatal("Timeout: numbers not generated")
			}
		}
	})

	t.Run("Channel closure on context cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		ch := RandomGenerator(ctx)

		time.Sleep(10 * time.Millisecond)

		cancel()

		select {
		case _, ok := <-ch:
			if !ok {
				t.Error("Channel should be closed")
			}
		case <-time.After(1 * time.Second):
			t.Error("Channel was not closed")
		}
	})

	t.Run("No data after context timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()

		ch := RandomGenerator(ctx)

		<-ctx.Done()

		select {
		case _, ok := <-ch:
			if !ok {
				t.Error("Channel should be closed after timeout")
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Channel should be closed")
		}
	})
}
