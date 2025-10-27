// Работа с массивами
package main

import (
	"fmt"
)

// Дан массив [1, 2, 3, 4, 5], найти сумму элементов
func plus(arr [5]int) {
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		fmt.Println(sum)

	}
}

// Дан массив [8, 3, 12, 5, 9], найти наибольший и наименьший элементы
func findMinMax(arr [5]int) (int, int) {
	max := arr[0]
	min := arr[0]

	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min
}

// Дан массив [2, 7, 4, 9, 6, 11], посчитать сколько четных чисел
func findEvenNums(arr []int) int {
	count := 0

	for _, nums := range arr {
		if nums%2 == 0 {
			count++
		}
	}
	return count
}

// Дан массив [10, 20, 30, 40, 50] и число 30, проверить его наличие
func findtargetNum(arr []int, target int) bool {
	a := target

	for _, i := range arr {
		if i == a {
			return true
		}
	}
	return false
}

// Дан массив [15, 8, 22, 5, 18, 9], найти все числа больше 10

func findNums(arr []int, target int) []int {
	res := []int{}

	for _, i := range arr {
		if i > target {
			res = append(res, i)
		}
	}
	return res
}

// Дан массив [1, 2, 2, 3, 4, 4, 5], вернуть массив без дубликатов
func delDuplicate(arr []int) []int {
	seen := make(map[int]bool)
	res := []int{}

	for _, i := range arr {
		if !seen[i] {
			seen[i] = true
			res = append(res, i)
		}
	}
	return res

}

// Дан массив [1, 2, 3, 4, 5], вернуть массив с удвоенными значениями
func double(arr []int) []int {
	result := []int{}

	for i := 0; i < len(arr); i++ {
		doubler := arr[i] + arr[i]

		result = append(result, doubler)
	}
	return result

}

// Дан массив [1, 2, 3, 4, 5], вернуть [5, 4, 3, 2, 1]

func rev(arr []int) []int {

	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - 1 - i
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
	return arr

}

// Даны массивы [1, 2, 3] и [4, 5, 6], объединить их

func arrplus(arr1 []int, arr2 []int) []int {
	result := append(arr1, arr2...)
	return result

}

// Даны массивы [1, 2, 3, 4] и [3, 4, 5, 6], найти общие элементы

func findboth(arr1 []int, arr2 []int) []int {
	both := []int{}

	for _, value1 := range arr1 {
		for _, value2 := range arr2 {
			if value1 == value2 {
				both = append(both, value1)
			}
		}

	}
	return both

}

func main() {
	fmt.Println("=== ПРОВЕРКА ВСЕХ ФУНКЦИЙ ===")

	// 1. Поиск четных чисел
	fmt.Println("\n1. Поиск колличества четных чисел:")
	arr1 := []int{2, 7, 4, 9, 6, 11}
	evenCount := findEvenNums(arr1)
	fmt.Printf("Массив: %v, Четных чисел: %d\n", arr1, evenCount)

	// 2. Поиск числа в массиве
	fmt.Println("\n2. Поиск числа в массиве:")
	arr2 := []int{10, 20, 30, 40, 50}
	target := 30
	found := findNums(arr2, target)
	fmt.Printf("Массив: %v, Число %d найдено: %d\n", arr2, target, found)

	// 3. Числа больше заданного
	fmt.Println("\n3. Числа больше заданного:")
	arr3 := []int{15, 8, 22, 5, 18, 9}
	numbers := findNums(arr3, 10)
	fmt.Printf("Массив: %v, Числа > 10: %v\n", arr3, numbers)

	// 4. Удаление дубликатов
	fmt.Println("\n4. Удаление дубликатов:")
	arr4 := []int{1, 2, 2, 3, 4, 4, 5}
	unique := delDuplicate(arr4)
	fmt.Printf("Было: %v, Стало: %v\n", arr4, unique)

	// 5. Удвоение элементов
	fmt.Println("\n5. Удвоение элементов:")
	arr5 := []int{1, 2, 3, 4, 5}
	doubled := double(arr5)
	fmt.Printf("Было: %v, Удвоено: %v\n", arr5, doubled)

	// 6. Переворот массива
	fmt.Println("\n6. Переворот массива:")
	arr6 := []int{1, 2, 3, 4, 5}
	reversed := rev(arr6)
	fmt.Printf("Было: %v, Перевернуто: %v\n", arr6, reversed)

	// 7. Объединение массивов
	fmt.Println("\n7. Объединение массивов:")
	arr7a := []int{1, 2, 3}
	arr7b := []int{4, 5, 6}
	combined := arrplus(arr7a, arr7b)
	fmt.Printf("Массив 1: %v, Массив 2: %v, Объединенный: %v\n", arr7a, arr7b, combined)

	// 8. Общие элементы
	fmt.Println("\n8. Общие элементы двух массивов:")
	arr8a := []int{1, 2, 3, 4}
	arr8b := []int{3, 4, 5, 6}
	common := findboth(arr8a, arr8b)
	fmt.Printf("Массив 1: %v, Массив 2: %v, Общие: %v\n", arr8a, arr8b, common)

	fmt.Println("\n=== ВСЕ ФУНКЦИИ ПРОВЕРЕНЫ ===")
}
