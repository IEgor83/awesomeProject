package main

import "fmt"

func gorutine_test(channel chan string) {
	fmt.Println("Hey, " + <-channel + "!")
}

func main() {
	fmt.Println("main() started")
	channel := make(chan string)

	go gorutine_test(channel)

	channel <- "Rob"
	fmt.Println("main() stopped")
}
