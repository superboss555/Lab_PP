package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Data структура для обработки JSON данных
type Data struct {
	Message string `json:"message"` // Используйте обратные кавычки для указания тегов
}

// helloHandler обрабатывает GET запросы на /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!")
}

// dataHandler обрабатывает POST запросы на /data
func dataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data Data
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Выводим данные в консоль
	fmt.Printf("Received data: %s\n", data.Message)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data received: %s", data.Message)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/data", dataHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
