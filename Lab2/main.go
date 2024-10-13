package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {

	// task 1
	fmt.Println("\n\n task 1:")
	var number int
	fmt.Print("Введите число: ")
	fmt.Fscan(os.Stdin, &number)

	if number%2 == 0 {
		fmt.Println("чётное")
	} else {
		fmt.Println("нечётное")
	}

	//task 2
	fmt.Println("\n\n task 2:")
	numb(number)

	//task 3
	fmt.Println("\n\n task 3:")
	for i := 1; i < 10; i++ {
		fmt.Print(i, "	")
	}

	//task 4
	fmt.Println("\n\n task 4:")
	var str string
	fmt.Print("Введите строку: ")
	fmt.Fscan(os.Stdin, &str)
	var length int = 0
	length = strLength(str, length)
	fmt.Print("длина строки = ", length)

	//task 5
	fmt.Println("\n\n task 5:")
	rect := Rectangle{
		Length: 5.0,
		Width:  3.0,
	}
	area := rect.Area()
	fmt.Printf("Площадь прямоугольника: %.2f\n", area)

	//task 6
	fmt.Println("\n\n task 6:")
	fmt.Print("среднее значение чисел ", number, " ", length, " = ", average(number, length))
}

func numb(x int) {
	if x > 0 {
		fmt.Println("Positive")
	} else if x < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
}

func strLength(str string, length int) int {
	length = utf8.RuneCountInString(str)
	return length
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func average(a int, b int) int {
	return int(a+b) / 2
}
