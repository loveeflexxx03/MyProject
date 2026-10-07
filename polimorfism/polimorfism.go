package main

import (
	"fmt"
	"math"
)

// Shape — интерфейс для геометрических фигур
type Shape interface {
	Area() float64
}

// Rectangle — прямоугольник
type Rectangle struct {
	Width  float64
	Height float64
}

// Area реализует интерфейс Shape для Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle — круг
type Circle struct {
	Radius float64
}

// Area реализует интерфейс Shape для Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 2.5},
		Rectangle{Width: 5, Height: 5},
		Circle{Radius: 1},
	}

	var total float64
	for _, s := range shapes {
		fmt.Println("Площадь:", s.Area())
		total += s.Area()
	}
	fmt.Println("Общая площадь:", total)
}
