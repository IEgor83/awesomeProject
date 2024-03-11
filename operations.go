package main

import "fmt"

func main() {
	// Создание переменных разных типов
	number := 10
	var pi float64 = 3.14
	var name string = "Alice"
	var isTrue bool = true

	fmt.Printf("Имя: %s, Число пи:  %.2f\n", name, pi)

	// Конструкция if-else
	if isTrue {
		fmt.Println("Условие истинно")
	} else {
		fmt.Println("Условие ложно")
	}
	// Цикл for
	for i := 0; i < number; i++ {
		fmt.Println("Итерация:", i)
	}
	// Массив
	var numbers [5]int
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i * 2
	}
	fmt.Println("Массив numbers:", numbers)
}
