package main

import (
	"fmt"
	"time"
)

func mine(transferPoint chan int, n int) {
	fmt.Println("Поход в шахту номер", n, "начался...")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту номер", n, "закончился")

	transferPoint <- 10
	fmt.Println("Рабочий уголь передал")

}

func main() {

	transferPoint := make(chan int, 3)

	coal := 0

	initTime := time.Now()

	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	coal += <-transferPoint
	fmt.Println("Поток забрал 1 раз")
	coal += <-transferPoint
	fmt.Println("Поток забрал 2 раз")
	coal += <-transferPoint
	fmt.Println("Поток забрал 3 раз")

	fmt.Println("Добыли:", coal)
	fmt.Println("Прошло времени:", time.Since(initTime))

}
