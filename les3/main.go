// Проверьте, является ли число палиндромом
// Палиндром читается одинаково слева направо и справа налево
// Пример: 121 -> true, -121 -> false, 10 -> false
package main

import "fmt"

// Напишите функцию, которая возвращает сумму двух чисел

func sum(a int, b int) int {
	sum := a + b
	return sum

}

// Напишите функцию, которая возвращает большее из двух чисел
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}

// Напишите функцию, которая возвращает true если число четное
func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

// Напишите функцию, которая возвращает сумму всех элементов массива
func factorial(n int) int {

	if n == 0 {
		return 1
	}

	result := 1
	for i := 1; i <= n; i++ {
		result = result * i

	}
	return result
}

// Напишите функцию, которая возвращает сумму всех элементов массива
func sumArray(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum = sum + i
	}
	return sum
}

// Напишите функцию, которая проверяет, есть ли число в массиве
func contains(arr []int, target int) bool {
	for i := range arr {
		if arr[i] == target {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(sumArray([]int{1, 2, 3}))    // 6
	fmt.Println(contains([]int{1, 2, 3}, 2)) // true
	fmt.Println(contains([]int{1, 2, 5}, 5)) // false
	fmt.Println(contains([]int{5, 2}, 1))    // false

}
