package main

import (
	"context"
	"math/rand"
	"time"
)

func RandomGenerator(ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		for {
			select {
			case <-ctx.Done():
				return
			default:
				out <- r.Intn(100)
			}
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for num := range RandomGenerator(ctx) {
		println(num)
	}
}
