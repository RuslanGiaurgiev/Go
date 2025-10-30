package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	dish int
	mu   sync.Mutex
)

func cook(wg *sync.WaitGroup, n int, ch chan<- int) {
	defer wg.Done()
	for {
		fmt.Println("Я повар", n, "начинаю кук")
		time.Sleep(2 * time.Second)
		fmt.Println("Диш готов")

		ch <- +1
		fmt.Println("Диш передан на стол заказов")

		mu.Lock()
		dish++
		mu.Unlock()

	}
}

func waiter(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	fmt.Println("Я соло официант начал работу(образно)")
	for {
		cookId := <-ch
		fmt.Println("Отнёс диш от повара:", cookId)
	}

}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	go cook(wg, 1, ch)
	go cook(wg, 2, ch)

	wg.Add(1)
	go waiter(wg, ch)

	time.Sleep(15 * time.Second)
	fmt.Println("Увы время вышло")

	fmt.Println("Суммарно вышло", dish)

}
