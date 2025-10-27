package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (p Bank) Pay(usd int) int {
	fmt.Println("Оплата банковским счётом")
	fmt.Println("Сумма оплаты:", usd, "$")

	id := rand.Int()

	return id
}

func (p Bank) Cancel(id int) {
	fmt.Println("Отмена операции через счёт №", id)

}
