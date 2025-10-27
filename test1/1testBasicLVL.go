// будем чуть чтуь разбираться с defer()
// и посмотрим на указатели

package main

import "fmt"

func main() {
	x := 5
	y := 7
	fmt.Println("До свапа:")
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	swap(&x, &y)
	fmt.Println("После свапа:")
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	z := 5

	fmt.Println("число до=", z)
	increment(&z)
	fmt.Println("число после=", z)

	a := 5

	safeDereference(&a)
	fmt.Println("Число:", a)
	safeDereference(&a)
	a = 0
	fmt.Println("Число 0", a)

}

// Напишите функцию swap, которая принимает два указателя на целые числа и меняет их значения местами.

func swap(x *int, y *int) {
	num := *x
	*x = *y
	*y = num

}

func increment(z *int) {
	*z = *z + 1

}

func safeDereference(num *int) int {
	if num != nil {
		return *num
	}
	return 0
}
