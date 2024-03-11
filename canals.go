package main

import "fmt"

func main() {
	channel := make(chan float32)

	fmt.Printf("type of 'c' is %T\n", channel)
	fmt.Printf("value of 'c' is %v\n", channel)
}
