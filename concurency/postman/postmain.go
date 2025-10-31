package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func postman(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- string, n int, mail string) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я почтальон:", n, "мой рабочий день окончен")
			return

		default:
			fmt.Println("Я почтальон:", n, "взял письмо")
			time.Sleep(1 * time.Second)
			fmt.Println("Я почтальон:", n, "принёс письмо:", mail)

			transferPoint <- mail

			fmt.Println("Я почтальон:", n, "положил в ящик письмо:", mail)
		}
	}

}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go postman(ctx, wg, mailTransferPoint, i, postmanMail(i))
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()
	return mailTransferPoint
}

func postmanMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "Новости",
		2: "Приглашение",
		3: "Информация",
	}
	mail, ok := ptm[postmanNumber]
	if !ok {
		return "Буклет"
	}

	return mail

}
