package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Определим структуру для математического запроса
type Request struct {
	Operation string
	A         float64
	B         float64
	Result    chan float64
}

// Калькулятор, который будет обрабатывать запросы
func calculator(requests <-chan Request, wg *sync.WaitGroup) {
	defer wg.Done()
	for req := range requests {
		var result float64
		switch req.Operation {
		case "+":
			result = req.A + req.B
		case "-":
			result = req.A - req.B
		case "*":
			result = req.A * req.B
		case "/":
			if req.B != 0 {
				result = req.A / req.B
			} else {
				log.Println("Ошибка: Деление на ноль")
				result = 0
			}
		default:
			log.Println("Ошибка: Неизвестная операция")
		}
		req.Result <- result // Отправка результата обратно через канал
	}
}

// Клиент, который отправляет запросы
func client(op string, a float64, b float64, requests chan<- Request) {
	resultChan := make(chan float64) // Канал для получения результата
	request := Request{op, a, b, resultChan}

	requests <- request    // Отправляем запрос на обработку
	result := <-resultChan // Получаем результат
	fmt.Printf("Запрос: %f %s %f = %f\n", a, op, b, result)
}

func main() {
	requests := make(chan Request)
	var wg sync.WaitGroup

	// Запустим горутину калькулятора
	wg.Add(1)
	go calculator(requests, &wg)

	// Случайные запросы от клиентских горутин
	operations := []string{"+", "-", "*", "/"}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		op := operations[rand.Intn(len(operations))]
		a := rand.Float64() * 100
		b := rand.Float64() * 100
		go client(op, a, b, requests)
		time.Sleep(100 * time.Millisecond) // Задержка между запросами
	}

	time.Sleep(2 * time.Second) // Даем время горутинам завершить выполнение

	close(requests) // Закрываем канал запросов, чтобы завершить калькулятор
	wg.Wait()       // Ждем завершения всех горутин
}
