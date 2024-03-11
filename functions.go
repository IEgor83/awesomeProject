package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	fmt.Println(f(), f(), f(), f(), f())

	result := add(5, 3)
	fmt.Println("Результат сложения:", result)
}
