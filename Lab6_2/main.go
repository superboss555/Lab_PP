package main

import (
	"fmt"
)

// Функция для генерации чисел Фибоначчи
func fibonacci(n int, ch chan<- int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x       // Отправляем текущее число в канал
		x, y = y, x+y // Генерируем следующее число Фибоначчи
	}
	close(ch) // Закрываем канал после завершения генерации
}

func main() {
	ch := make(chan int) // Создаем канал

	go fibonacci(15, ch) // Запускаем горутину для генерации чисел Фибоначчи

	// Чтение и вывод данных из канала
	for num := range ch { // Используем range для чтения из канала до его закрытия
		fmt.Println(num)
	}
}
