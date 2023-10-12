package shapes

type Triangle struct {
	BaseShape // Является композицией так как все используют повторно не вводя какие либо измнения
	Height    float64
	Base      float64
}

// Функция конструктор собирающая новую фигуру
func NewTriangle(name string, height, base float64) *Triangle {
	return &Triangle{
		BaseShape: BaseShape{Name: name},
		Height:    height,
		Base:      base,
	}
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.Height * t.Base
}
