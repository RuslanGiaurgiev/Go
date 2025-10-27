package main

import (
	"study/payments"
	"study/payments/methods"

	"github.com/k0kubun/pp/v3"
)

func main() {
	method := methods.NewBank()

	paymentModule := payments.NewPaymentMethod(method)

	idBurger := paymentModule.Pay("Burger", 5)
	paymentModule.Pay("Steam Game", 500)
	paymentModule.Cancel(idBurger)
	allInfo := paymentModule.AllInfo()

	pp.Println(allInfo)

}
