package main

import (
	"fmt"
	"time"
)

func main() {

	//задание 1 - определить время
	fmt.Println("Задание 1: определить текущее время и дату")
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("\n", "Текущее время: ", currentTime)

	//задание 2 - создать переменные разных типов

	fmt.Println("\n", "Задание 2: создать переменные разных типов")
	var (
		variable1 int     = 10
		variable2 float64 = 3.2345
		variable3 string  = "PP"
		variable4 bool    = true
	)

	fmt.Println("\n", "Переменная int: ", variable1, "\n",
		"Переменная float: ", variable2, "\n",
		"Переменная string: ", variable3, "\n",
		"Переменная bool: ", variable4)

	//задание 3 - использовать краткое определение переменной
	name := "PP"
	fmt.Println("\n", "Задание 3: использовать краткое определение переменной", "\n",
		"\n", "Создание переменной через краткое определение: ", name)

	//задание 4 -  выполнениe арифметических операций с двумя целыми числами

	fmt.Println("\n", "Задание 4:  выполнениe арифметических операций с двумя целыми числами")
	var variable5 int = 12
	var sum1 int = variable1 + variable5
	fmt.Println("\n", "Сумма двух целых чисел: ", sum1)

	//задание 5 - сумма и разность двух чисел с плавающей точкой

	fmt.Println("\n", "Задание 5: сумма и разность двух чисел с плавающей точкой")
	var variable6 float64 = 0.4352
	var sum2 float64 = variable2 + variable6
	var dif1 float64 = variable2 - variable6
	fmt.Println("\n", "Сумма двух чисел: ", sum2,
		"\n", "Разность двух чисел: ", dif1)

	//задание 6 - среднее значение трех чисел
	fmt.Println("\n", "Задание 6: среднее значение трех чисел",
		"\n", (variable2+variable6+sum2)/3)
}
