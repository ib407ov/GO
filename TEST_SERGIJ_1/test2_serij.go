package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestBody struct {
	Text   string `json:"input"`
	Header string `json:"data"`
}

func main() {
	http.HandleFunc("/", helloWorld)
	// Обробник для шляху "/data"
	http.HandleFunc("/data", handleRequest)

	// Запускаємо сервер на порту 8080
	http.ListenAndServe(":4000", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("HelloWorld")
	fmt.Fprintf(w, "Дані отримано успішно")
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Парсимо JSON-дані з тіла запиту
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Неправильний формат JSON", http.StatusBadRequest)
		return
	}

	// Виводимо дані у консоль
	fmt.Println(requestBody.Header)
	fmt.Println(requestBody.Text)

	// Відправляємо відповідь клієнту
	fmt.Fprintf(w, "Дані отримано успішно go naxuj")
}
