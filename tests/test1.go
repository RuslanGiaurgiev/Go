package main

import "fmt"

// Напиши функцию, которая принимает год рождения и возвращает возраст
func calculateAge(birthYear int) int {
	year := 2025
	age := year - birthYear
	return age

}

// Напиши функцию, которая проверяет, достаточно ли длинный пароль
// Пароль должен быть не менее 8 символов
func isPasswordStrong(password string) bool {
	for i := range len(password) {
		if i > 7 {
			return true
		}
	}
	return false

}

// Напиши функцию, которая считает сумму всех цен в списке
func totalPrice(prices []int) int {
	sum := 0
	for i := 0; i < len(prices); i++ {
		sum = sum + prices[i]
	}
	return sum

}
func totalPrice1(prices []int) int {
	sum := 0
	for _, price := range prices {

		sum = sum + price
	}
	return sum
}

// Напиши функцию, которая считает сколько раз число встречается в списке
func countOccurrences(numbers []int, target int) int {
	num := 0
	for _, i := range numbers {
		if i == target {
			num++
		}
	}
	return num

}

// Напиши функцию, которая переворачивает список
func reverseList(numbers []int) []int {
	rev := []int{}
	for i := len(numbers) - 1; i >= 0; i-- {
		rev = append(rev, numbers[i])
	}
	return rev
}

// Напиши функцию, которая находит самое маленькое число в списке
func findMin(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if min > numbers[i] {
			min = numbers[i]
		}
	}
	return min
}
func main() {
	fmt.Println("Ваш возраст: ", calculateAge(1984))
	fmt.Println(isPasswordStrong("555555558"))
	fmt.Println("=== ТЕСТ 1 ===")
	list1 := []int{100, 200, 300}
	result1 := totalPrice1(list1)
	fmt.Printf("Результат: %d\n", result1)
	fmt.Println(countOccurrences([]int{1, 2, 3, 2, 4, 2}, 2)) // 3
	fmt.Println("=== ТЕСТ 4 ===")
	list4 := []int{10, 20, 30}
	result4 := reverseList(list4)
	fmt.Printf("Результат: %v\n", result4)
	fmt.Println("=== ТЕСТ 5 ===")
	list2 := []int{5, 2, 8, 7, 9, 0, 1, -5}
	result2 := findMin(list2)
	fmt.Printf("Результат: %d\n\n", result2)
}
