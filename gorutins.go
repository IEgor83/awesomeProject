package main

import (
	"fmt"
	"time"
)

// Функция, которая будет запущена в горутине
func printNumbers1() {
	for i := 0; i < 5; i++ {
		fmt.Println("Горутина 1:", i)
		time.Sleep(time.Millisecond * 500) // Имитация задержки
	}
}

func printNumbers2() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Горутина: ", i)
	}
}

// Функция, которая будет запущена во второй горутине
func printLetters() {
	for char := 'a'; char < 'e'; char++ {
		fmt.Println("Горутина 2:", string(char))
		time.Sleep(time.Millisecond * 700) // Имитация задержки
	}
}

/*func main() {
	go printNumbers1()
	for i := 1; i <= 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("Основная горутина: ", i)
	}
}*/

func main() {
	go printNumbers1()
	go printLetters()

	for i := 0; i < 5; i++ {
		fmt.Println("Главная горутина:", i)
		time.Sleep(time.Millisecond * 600)
	}

	time.Sleep(time.Second * 3)
}
