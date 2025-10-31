package main

import "fmt"

//1. "Синхронизатор"
//Задача: Создай 3 горутины, которые печатают "A", "B", "C".
//Сделай так, чтобы они всегда выполнялись в порядке A → B → C, независимо от того в каком порядке запускаются.

func synco(ch1 chan string, ch2 chan string, ch3 chan string) {
	for id := 1; id <= 3; id++ {
		fmt.Println("Я горутина", id, "и я вывожу", letter)
		if id == 1 {
			ch1 <- "A"
		}
		if id == 2 {
			ch2 <- "B"
		}
		if id == 3 {
			ch3 <- "C"
		}
	}
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go synco(ch1, ch2, ch3)
	go synco(ch1, ch2, ch3)
	go synco(ch1, ch2, ch3)
}
