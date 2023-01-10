package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) { // Создаем контроллер
	w.Write([]byte("Hello, world!\n")) // Содержимое страницы
}

func main() {
	http.HandleFunc("/", Handler) // Задаем путь станицы с содержимым

	log.Println("Server starts on 8081 port...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
