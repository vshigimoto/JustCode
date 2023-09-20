package shapes

type Rectangle struct {
	BaseShape // Является композицией так как все используют повторно не вводя какие либо измнения
	Width     float64
	Height    float64
}

// Функция конструктор собирающая новую фигуру
func NewRectangle(name string, width, height float64) *Rectangle {
	return &Rectangle{
		BaseShape: BaseShape{Name: name},
		Width:     width,
		Height:    height,
	}
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}
