package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//	func main() {
//		http.HandleFunc("/", helloWorld)
//		//	http.HandleFunc("/", writePost)
//		port := 3000
//		fmt.Printf("Server is running on port %d\n", port)
//		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
//	}
func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("HelloWorld")
	fmt.Fprintf(w, "Дані отримано успішно")
}

type RequestBody struct {
	Text   string `json:"input"`
	Header string `json:"data"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	//// Отримуємо тіло POST-запиту
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, "Не вдалося прочитати тіло запиту", http.StatusBadRequest)
	//	return
	//}

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

func main() {
	http.HandleFunc("/", helloWorld)
	// Обробник для шляху "/data"
	http.HandleFunc("/data", handleRequest)

	// Запускаємо сервер на порту 8080
	http.ListenAndServe(":4000", nil)
}
