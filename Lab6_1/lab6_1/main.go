package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Функция для вычисления факториала
func factorial(n int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик после завершения
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		time.Sleep(100 * time.Millisecond) // Имитируем задержку
	}
	fmt.Printf("Факториал %d: %d\n", n, result)
}

// Функция для генерации случайных чисел
func generateRandomNumbers(count int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик после завершения
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(100)        // Генерируем случайное число от 0 до 99
		time.Sleep(150 * time.Millisecond) // Имитируем задержку
	}
	fmt.Printf("Случайные числа: %v\n", numbers)
}

// Функция для вычисления суммы числового ряда
func sumOfSeries(n int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик после завершения
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		time.Sleep(50 * time.Millisecond) // Имитируем задержку
	}
	fmt.Printf("Сумма числового ряда от 1 до %d: %d\n", n, sum)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3) // Устанавливаем количество горутин

	// Запускаем горутины
	go factorial(5, &wg)             // Расчет факториала числа 5
	go generateRandomNumbers(5, &wg) // Генерация 5 случайных чисел
	go sumOfSeries(10, &wg)          // Сумма ряда от 1 до 10

	// Ждем завершения всех горутин
	wg.Wait()
}
