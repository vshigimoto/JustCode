package shapes

import "fmt"

type Shape interface {
	Area() float64 // Является базовой функции которую все наследуют и модифицируют
	Display()      // функция для отображения
}

type BaseShape struct {
	Name string
}

// функция для отображения фигуры
func (s *BaseShape) Display() {
	fmt.Printf("Shape: %s\n", s.Name)
}
