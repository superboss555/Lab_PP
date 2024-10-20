package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Data структура для обработки JSON данных
type Data struct {
	Message string `json:"message"`
}

// helloHandler обрабатывает GET запросы на /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!")
}

// dataHandler обрабатывает POST запросы на /data
func dataHandler(w http.ResponseWriter, r *http.Request) {
	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Выводим данные в консоль
	fmt.Printf("Received data: %s\n", data.Message)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data received: %s", data.Message)
}

// loggingMiddleware логирует входящие запросы
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Передаем управление следующему обработчику
		next.ServeHTTP(w, r)

		// Логируем информацию о запросе
		log.Printf(
			"%s %s took %s",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/data", dataHandler)

	// Оборачиваем маршрутизатор в middleware
	loggedMux := loggingMiddleware(mux)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
