package main

import "fmt"

func main() {
	person := Person{
		Name: "Артурик",
		Age:  25,
	}

	acc := Account{
		Owner:   "Артурик",
		Balance: 150,
	}

	fmt.Println("До:", person)
	person.birthday()
	fmt.Println("После:", person)

	fmt.Println("До:")
	fmt.Println(acc)
	acc.deposit(50)
	fmt.Println("После:", acc)

}

// Cоздайте структуру Person с полями Name (string) и Age (int). Напишите метод birthday, который увеличивает возраст на 1.
type Person struct {
	Name string
	Age  int
}

func (p *Person) birthday() int {
	p.Age++
	return p.Age

}

//Создайте структуру Account с полями Owner (string) и Balance (float64).
// Напишите функцию deposit, которая принимает указатель на Account и сумму для пополнения.

type Account struct {
	Owner   string
	Balance float64
}

func (a *Account) deposit(dep float64) float64 {
	a.Balance = a.Balance + dep
	return a.Balance

}

//Напишите функцию resetPerson, которая обнуляет все поля переданной структуры Person.

func (p *Person) resetPerson() int {
	res := p.Age - p.birthday()

	return res
}
