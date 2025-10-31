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
			fmt.Println("Foo звершил работу")
			return
		default:
			fmt.Println("foo")
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func boo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Boo звершил работу")
			return
		default:
			fmt.Println("boo")
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func main() {

	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)

	go foo(parentContext)
	go boo(childContext)

	time.Sleep(2 * time.Second)
	childCancel()

	time.Sleep(2 * time.Second)
	parentCancel()

	time.Sleep(2 * time.Second)
	fmt.Println("Все горутины закончились")
}
