package main

import (
	f "fmt"
	"strings"
)

func main() {

	//task 1
	f.Println("\n\n task 1: ")
	ages := make(map[string]int)

	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35

	f.Println("Текущие записи:")
	for name, age := range ages {
		f.Printf("%s: %d лет\n", name, age)
	}

	ages["David"] = 28

	f.Println("\nЗаписи после добавления нового человека:")
	for name, age := range ages {
		f.Printf("%s: %d лет\n", name, age)
	}

	//task 2
	f.Println("\n\n task 2: ")
	avg := averageAge(ages)
	f.Printf("Средний возраст: %.2f лет\n", avg)

	//task 3
	f.Println("\n\n task 3: ")
	var nameToDelete string
	f.Print("\nВведите имя для удаления: ")
	f.Scanln(&nameToDelete)
	deletePerson(ages, nameToDelete)
	f.Println("\nОбновленные записи:")
	for name, age := range ages {
		f.Printf("%s: %d лет\n", name, age)
	}

	//task 4
	f.Println("\n\n task 4: ")
	f.Print("Введите строку: ")
	var input string
	f.Scanln(&input)

	upperCaseInput := strings.ToUpper(input)

	f.Println("Верхний регистр:", upperCaseInput)

	//task 5
	f.Println("\n\n task 5: ")
	var sum int
	var number int

	f.Println("Введите числа (введите '0' для завершения):")

	for {
		f.Print("Введите число: ")
		f.Scan(&number)

		if number == 0 {
			break
		}

		sum += number
	}

	f.Printf("Сумма введенных чисел: %d\n", sum)

	//task 6
	f.Println("\n\n task 6: ")
	var n int

	f.Print("Введите количество чисел: ")
	f.Scan(&n)

	numbers := make([]int, n)

	f.Println("Введите числа:")
	for i := 0; i < n; i++ {
		f.Print("Число ", i+1, ": ")
		f.Scan(&numbers[i])
	}

	f.Println("Числа в обратном порядке:")
	for i := n - 1; i >= 0; i-- {
		f.Println(numbers[i])
	}
}

func averageAge(ages map[string]int) float64 {
	var total int
	var count int

	for _, age := range ages {
		total += age
		count++
	}

	if count == 0 {
		return 0
	}

	return float64(total) / float64(count)
}

func deletePerson(ages map[string]int, name string) {
	if _, exists := ages[name]; exists {
		delete(ages, name)
		f.Printf("Запись с именем '%s' удалена.\n", name)
	} else {
		f.Printf("Запись с именем '%s' не найдена.\n", name)
	}
}
