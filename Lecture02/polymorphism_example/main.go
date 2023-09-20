package main

import (
	"fmt"
	"polymorphism_example/shapes" // импорт фигур
)

func main() {
	circle := shapes.NewCircle("Circle 1", 5.0)               // создание новой фигуры и задача ее свойств
	rectangle := shapes.NewRectangle("Rectangle 1", 4.0, 6.0) // создание новой фигуры и задача ее свойств
	triangle := shapes.NewTriangle("Triangle 1", 10, 5)
	shapesList := []shapes.Shape{circle, rectangle, triangle} // создание массива фигур для их дальнейшего вывода

	for _, shape := range shapesList {
		shape.Display()                          // Вывод название и тип фигуры
		fmt.Printf("Area: %.2f\n", shape.Area()) // Вывод площади фигуры
	}
}
