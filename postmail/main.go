package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	totalLetters int
	mu           sync.Mutex
)

func postman(wg *sync.WaitGroup, n int, mailbox chan<- int) {

	defer wg.Done()

	for {
		fmt.Println("Почтальон", n, "пошёл доставлять письмо")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Почтальон", n, "ппришёл и готов отдать письмо")

		mailbox <- n
		fmt.Println("Почтальон", n, "передал письмо в ящик")

		mu.Lock()
		totalLetters++
		mu.Unlock()
	}

}

func sorter(wg *sync.WaitGroup, mailbox <-chan int) {
	defer wg.Done()

	for {
		postmanID := <-mailbox
		fmt.Println("Сортировщик обработал письмо от почтальона", postmanID)
		time.Sleep(500 * time.Millisecond)
	}

}

func main() {
	var wg sync.WaitGroup
	mailbox := make(chan int, 4)

	timeInit := time.Now()

	wg.Add(50)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)
	go sorter(&wg, mailbox)

	wg.Add(50)
	go postman(&wg, 1, mailbox)
	go postman(&wg, 2, mailbox)
	go postman(&wg, 3, mailbox)
	go postman(&wg, 4, mailbox)
	go postman(&wg, 5, mailbox)
	go postman(&wg, 1, mailbox)
	go postman(&wg, 2, mailbox)
	go postman(&wg, 3, mailbox)
	go postman(&wg, 4, mailbox)
	go postman(&wg, 5, mailbox)
	go postman(&wg, 1, mailbox)
	go postman(&wg, 2, mailbox)
	go postman(&wg, 3, mailbox)
	go postman(&wg, 4, mailbox)
	go postman(&wg, 5, mailbox)
	go postman(&wg, 1, mailbox)
	go postman(&wg, 2, mailbox)
	go postman(&wg, 3, mailbox)
	go postman(&wg, 4, mailbox)
	go postman(&wg, 5, mailbox)
	go postman(&wg, 1, mailbox)
	go postman(&wg, 2, mailbox)
	go postman(&wg, 3, mailbox)
	go postman(&wg, 4, mailbox)
	go postman(&wg, 5, mailbox)
	go postman(&wg, 1, mailbox)
	go postman(&wg, 2, mailbox)
	go postman(&wg, 3, mailbox)
	go postman(&wg, 4, mailbox)
	go postman(&wg, 5, mailbox)
	go postman(&wg, 1, mailbox)
	go postman(&wg, 2, mailbox)
	go postman(&wg, 3, mailbox)
	go postman(&wg, 4, mailbox)
	go postman(&wg, 5, mailbox)
	go postman(&wg, 1, mailbox)
	go postman(&wg, 2, mailbox)
	go postman(&wg, 3, mailbox)
	go postman(&wg, 4, mailbox)
	go postman(&wg, 5, mailbox)
	go postman(&wg, 1, mailbox)
	go postman(&wg, 2, mailbox)
	go postman(&wg, 3, mailbox)
	go postman(&wg, 4, mailbox)
	go postman(&wg, 5, mailbox)
	go postman(&wg, 1, mailbox)
	go postman(&wg, 2, mailbox)
	go postman(&wg, 3, mailbox)
	go postman(&wg, 4, mailbox)
	go postman(&wg, 5, mailbox)

	time.Sleep(10 * time.Second)

	fmt.Println("Время вышло")
	fmt.Println("Всего доставлено писем:", totalLetters)
	fmt.Println("Времени заняло:", time.Since(timeInit))
	close(mailbox)
}
