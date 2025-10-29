package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Foo закончился")
			return
		default:
			fmt.Println("Foo продолжается")
		}

		time.Sleep(100 * time.Millisecond)
	}

}

func boo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Boo закончился")
			return
		default:
			fmt.Println("Boo продолжается")
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	firstcontext, firstcancel := context.WithCancel(context.Background())
	secondcontext, secondcancel := context.WithCancel(firstcontext)

	go foo(firstcontext)
	go boo(secondcontext)

	time.Sleep(1 * time.Second)
	secondcancel()

	time.Sleep(1 * time.Second)
	firstcancel()

	time.Sleep(3 * time.Second)
}
