package main

import (
	"concurency/miner"
	"concurency/postman"
	"context"
	"fmt"
	"time"
)

func main() {
	var coal int
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(5 * time.Second)
		minerCancel()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		postmanCancel()
	}()

	coaltransferPoint := miner.MinerPool(minerContext, 50)
	mailTransferPoint := postman.PostmanPool(postmanContext, 50)

	isCoalClosed := false
	isMailClosed := false

	for !isCoalClosed || !isMailClosed {
		select {
		case c, ok := <-coaltransferPoint:
			if !ok {
				isCoalClosed = true
				continue
			}

			coal += c

		case p, ok := <-mailTransferPoint:
			if !ok {
				isMailClosed = true
				continue
			}

			mails = append(mails, p)

		}

	}
	fmt.Println("Суммарно получено угля:", coal)
	fmt.Println("Суииарно получено писем:", len(mails))

}
