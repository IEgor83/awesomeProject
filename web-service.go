package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to our website!</h1>")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About Us</h1><p>This is a simple Go web application.</p>")
}

func main() {
	// Регистрируем обработчики для различных URL-адресов
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)

	// Запускаем веб-сервер на порту 8080
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
