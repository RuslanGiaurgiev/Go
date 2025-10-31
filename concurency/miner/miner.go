package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n int,
	power int,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я шахтёр номер", n, "закончил работу")
			return
		default:
			fmt.Println("Я шахтёр номер", n, "начал работу", power)
			time.Sleep(1 * time.Second)
			fmt.Println("Я шахтёр:", n, "добыл уголь", power)

			transferPoint <- power
			fmt.Println("Я шахтёр:", n, "передал уголь", power)

		}

	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	CoaltransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 1; i < minerCount; i++ {
		wg.Add(1)
		go Miner(ctx, wg, CoaltransferPoint, i, i*5)
	}

	go func() {
		wg.Wait()
		close(CoaltransferPoint)
	}()

	return CoaltransferPoint

}
