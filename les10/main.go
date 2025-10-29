package main

import (
	"fmt"
	"sync"
	"time"
)

func postmain(wg *sync.WaitGroup, text string) {
	for i := 1; i <= 3; i++ {
		fmt.Println("Я отнёс газету:", text, "B", i, "раз")
		time.Sleep(500 * time.Millisecond)

	}

	wg.Done()

}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go postmain(wg, "Новости")

	wg.Add(1)
	go postmain(wg, "Программы")

	wg.Add(1)
	go postmain(wg, "Игровой журнал")

	wg.Wait()

	time.Sleep(1 * time.Second)
	fmt.Println("Main завершился")
}
