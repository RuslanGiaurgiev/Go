// просто протестирую состяние гонки

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var number1 int = 0
var number2 atomic.Int64

func increase1(wg1 *sync.WaitGroup) {
	defer wg1.Done()

	for i := 0; i < 2000; i++ {
		number1++
	}
}

func increase2(wg2 *sync.WaitGroup) {
	defer wg2.Done()

	for i := 0; i < 2000; i++ {
		number2.Add(1)
	}
}

func main() {
	wg1 := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}

	wg1.Add(5)
	go increase1(wg1)
	go increase1(wg1)
	go increase1(wg1)
	go increase1(wg1)
	go increase1(wg1)

	wg1.Wait()
	fmt.Println("Число на исходе", number1)

	wg2.Add(5)
	go increase2(wg2)
	go increase2(wg2)
	go increase2(wg2)
	go increase2(wg2)
	go increase2(wg2)

	wg2.Wait()
	fmt.Println("Число на исходе", number2.Load())

}
