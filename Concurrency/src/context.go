package main

import (
	"context"
	"fmt"
	"time"
)

// START OMIT
func good(ctx context.Context) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
		default:
			fmt.Println("G", i)
			time.Sleep(1 * time.Second)
		}
	}
}
func bad() {
	for i := 0; ; i++ {
		fmt.Println("B", i)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go good(ctx)
	go bad()
	time.Sleep(5 * time.Second)
	cancel()
	fmt.Println("canceled")
	time.Sleep(2 * time.Second)
}

// END OMIT
