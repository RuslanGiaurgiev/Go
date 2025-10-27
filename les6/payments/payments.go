package payments

// разница между мапой и слайсом:
// мапа: O(1) получение выполняется быстрее т.к 1 константа, но невозможно получить последовательный массив то есть
// нельзя будет получить по какому нибудь порядку value
// слайс O(N), зависит от количества итераций, создаётся последовательный массив,
// да дольше по поиску, но есть возможность отфильтровать данные по порядку

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentsInfo  map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentMethod(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentsInfo:  make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

// поулчать информацию по покупке str, сумма покупки
// возвращает id операции
func (p *PaymentModule) Pay(description string, usd int) int {
	// проводим операцию и сохраняем в переменную id
	// id указывать не нужно, т.к мы используем его в мапе
	id := p.paymentMethod.Pay(usd)

	// вернуть описание, сумму и была ли операция отменена или нет?
	// вернуть мапу
	info := PaymentInfo{
		Description: description,
		USD:         usd,
		Cancelled:   false,
	}
	p.paymentsInfo[id] = info

	return id

}

// получает id покупки
// возвращает ничего
func (p *PaymentModule) Cancel(id int) {

	// проверяем оплату, если она не было совершена, но сразу завершаем функцию
	// в теории такого быть не может но доп проверка приветсвуется.
	info, ok := p.paymentsInfo[id]
	if !ok {
		return
	}
	// выполняем метод cansel по id операции
	p.paymentMethod.Cancel(id)

	// выстваляем cancelled который находится в info.go на true указывая тем самым что оперция была отменена.
	info.Cancelled = true
	p.paymentsInfo[id] = info

	// обязательно перезаписываем новую инфу в мапу id операции т.к выше мы получили только копию

}

// получает id, инфо (str), отменена ли? (bool)
func (p *PaymentModule) Info(id int) PaymentInfo {

	//опять же доп проеврка, мало ли инфы нету, если её нету, мы выходим функцию
	info, ok := p.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}
	// если всё ок, радуемся и получаем информацию об операции
	return info
}

// ничего не принимает
// возвращеает все операции
func (p *PaymentModule) AllInfo() map[int]PaymentInfo {

	// чтобы не дай Бог никто не заруинил нам основные данные создаём новую мапу, указывая длину элементов основной мапы,
	// чтобы комп не нагружать
	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))
	// и после пробегая по всем элементам основной мапы сохраняем их значения в ключ новой мапы.
	for k, v := range p.paymentsInfo {
		tempMap[k] = v
	}
	// возвращаем само собой новую мапу
	return tempMap

}
