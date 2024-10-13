package main

import (
	f "fmt"
)

func main() {

	//task 1-2
	f.Println("\n\n task 1-2")
	num := 5
	result := Factorial(num)
	if result == -1 {
		f.Println("Факториал не определен для отрицательных чисел.")
	} else {
		f.Printf("Факториал числа %d равен %d\n", num, result)
	}

	//task 3
	f.Println("\n\n task 3")
	str := "Привет, мир!"
	reversedStr := Reverse(str)
	f.Printf("Исходная строка: %s\n", str)
	f.Printf("Перевернутая строка: %s\n", reversedStr)

	//task 4
	f.Println("\n\n task 4")
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
		f.Print(arr[i], " ")
	}

	//task 5
	f.Println("\n\n task 5")
	slice := arr[:]

	f.Println("Исходный срез:", slice)

	// Добавление элемента в конец среза
	slice = append(slice, 50)
	f.Println("Срез после добавления элемента 50:", slice)

	// Добавление нескольких элементов
	slice = append(slice, 60, 70)
	f.Println("Срез после добавления элементов 60 и 70:", slice)

	// Удаление элемента (например, элемента с индексом 2)
	indexToRemove := 2
	slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	f.Println("Срез после удаления элемента с индексом 2:", slice)

	// Удаление последнего элемента
	slice = slice[:len(slice)-1] // Удаляем последний элемент
	f.Println("Срез после удаления последнего элемента:", slice)

	//task 6
	f.Println("\n\n task 6")
	strings := []string{"Go", "ab", "task", "program", "like", "sad"}

	var longestString string

	for _, str := range strings {
		if len(str) > len(longestString) {
			longestString = str
		}
	}

	// Выводим результат
	f.Println("Самая длинная строка:", longestString)
}
