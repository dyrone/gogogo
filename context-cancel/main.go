package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		cancel()
	}()

	for n := range gen(ctx) {
		fmt.Println(n)
		//	if n == 5 {
		//		break
		//	}
	}

}
