package shapes

import "math"

type Circle struct {
	BaseShape // Является композицией так как все используют повторно не вводя какие либо измнения
	Radius    float64
}

// Функция конструктор собирающая новую фигуру
func NewCircle(name string, radius float64) *Circle {
	return &Circle{
		BaseShape: BaseShape{Name: name},
		Radius:    radius,
	}
}

// В данном случае функция Area показывает полиморфизм так как дополняет базовый класс Area Shape
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
