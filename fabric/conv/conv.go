package conv

import (
	"context"
	"fmt"
	"time"
)

func Conveyor(ctx context.Context, transferpoint chan<- int, n int) {
	for {
		fmt.Println(n, "- Конвеер начал производить конфеты")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(n, "- Конвеер произвёл конфеты")
		transferpoint <- +1
		fmt.Println(n, "- Конвеер положил конфету")

	}
}

func ConveyorPool(ctx context.Context, ConveyorCount int) {
	Mainline := make(chan int, 5)

	go Conveyor(ctx, Mainline, 1)
	go Conveyor(ctx, Mainline, 2)

}

func Packager(Mainline chan int) {

}
