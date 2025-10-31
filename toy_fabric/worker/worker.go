package worker

import (
	"fmt"
	"sync"
	"time"
	"toy_fabric/controller"
)

type Worker struct {
	Id        int
	TotalMade int
}

var mu sync.Mutex

func (w *Worker) Worker(ch chan<- controller.Toy) {
	for i := 1; ; i++ {
		time.Sleep(1 * time.Second)

		toy := controller.Toy{
			Id:       i,
			Workerid: w.Id,
		}

		ch <- toy

		mu.Lock()
		w.TotalMade++
		mu.Unlock()

		fmt.Println("Рабочий", w.Id, "сделал и передал игрушку", i)

	}
}

func (w *Worker) WorkerStats() int {
	mu.Lock()
	defer mu.Unlock()
	return w.TotalMade

}
