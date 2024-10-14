package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers(ch chan<- int) {
	for i := 0; i < 15; i++ {
		num := rand.Intn(100)
		ch <- num
		time.Sleep(time.Second)
	}
	close(ch)
}

func checkEvenOdd(num int, ch chan<- string) {
	if num%2 == 0 {
		ch <- fmt.Sprintf("Число %d чётное", num)
	} else {
		ch <- fmt.Sprintf("Число %d нечётное", num)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Инициализируем генератор случайных чисел (устарел, но ещё работает)

	numbersCh := make(chan int)
	resultsCh := make(chan string)

	go generateRandomNumbers(numbersCh)

	done := 0 // Переменная для запуска завершения после 15 итераций

	for {
		select {
		case num, ok := <-numbersCh: // Принимаем случайное число из канала
			if !ok { // Проверяем, закрыт ли канал
				// Если канал закрыт, выходим из цикла
				if done < 15 {
					fmt.Println("Генерация завершена!")
				}
				return
			}
			go checkEvenOdd(num, resultsCh) // Запускаем проверку чётности в новой горутине
			done++
		case result := <-resultsCh:
			fmt.Println(result)
		}
	}
}
