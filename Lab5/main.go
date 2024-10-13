package main

import (
	f "fmt"
	"math"
)

//_______________________________________________________________
//task 1

type Person struct {
	Name string
	Age  int
}

func (p Person) Info() string {
	return f.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// _______________________________________________________________
// task 2
func (p *Person) Birthday() {
	p.Age++
}

// _______________________________________________________________
// task 3
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// _______________________________________________________________
// task 4
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	f.Printf("Площадь: %.2f\n", s.Area())
}

// _______________________________________________________________
// task 5
func printAreas(shapes []Shape) {
	for _, shape := range shapes {
		f.Printf("Площадь: %.2f\n", shape.Area())
	}
}

// _______________________________________________________________
// task 6
type Stringer interface {
	String() string
}

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) String() string {
	return f.Sprintf("Название: %s, Автор: %s, Год: %d", b.Title, b.Author, b.Year)
}

func main() {

	// task 1
	f.Println("\n\n task 1: ")
	person := Person{Name: "Alice", Age: 30}
	f.Println(person.Info())

	// task 2
	f.Println("\n\n task 2: ")
	person.Birthday()
	f.Println(person.Info())

	// task 3
	f.Println("\n\n task 3: ")
	circle := Circle{Radius: 5.0}

	area := circle.Area()
	f.Printf("Площадь круга с радиусом %.2f равна %.2f\n", circle.Radius, area)

	// task 4
	f.Println("\n\n task 4: ")
	rectangle := Rectangle{Width: 4.0, Height: 6.0}
	printArea(circle)
	printArea(rectangle)

	// task 5
	f.Println("\n\n task 5: ")
	shapes := []Shape{circle, rectangle}
	printAreas(shapes)

	// task 6
	f.Println("\n\n task 6: ")
	book := Book{
		Title:  "1984",
		Author: "Джордж Оруэлл",
		Year:   1949,
	}

	f.Println(book.String())

}
