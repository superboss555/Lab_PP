package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Worker - функция, которая будет обрабатывать задачи
func worker(id int, tasks <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		reversed := reverseString(task)                       // Реверсируем строку
		results <- fmt.Sprintf("Worker %d: %s", id, reversed) // Отправка результата в канал
	}
}

// reverseString - функция для реверсирования строки
func reverseString(s string) string {
	runes := []rune(s) // Работает с рунами для поддержки юникода
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// Чтение числа воркеров от пользователя
	var numWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&numWorkers)

	// Каналы для задач и результатов
	tasks := make(chan string, 100)
	results := make(chan string, 100)

	var wg sync.WaitGroup

	// Запускаем воркеры
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, tasks, results, &wg)
	}

	// Обработка ввода с клавиатуры в основной горутине
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите строки (для завершения введите пустую строку):")
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			if line == "ALL" {
				break // Завершаем ввод, если введена пустая строка
			}
			tasks <- line // Отправляем строку в канал задач
		}
		close(tasks) // Закрываем канал задач после ввода
	}()

	// Запускаем горутину для сбора результатов
	go func() {
		wg.Wait()      // Ждем завершения всех воркеров
		close(results) // Закрываем канал результатов
	}()

	// Вывод результатов
	for result := range results {
		fmt.Println(result) // Выводим результат на консоль
	}

	fmt.Println("Все воркеры завершили работу.")
}
