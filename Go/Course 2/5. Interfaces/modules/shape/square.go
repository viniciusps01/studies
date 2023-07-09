package shape

type Square struct {
	SideLength float64
}

func (s Square) GetArea() float64 {
	return s.SideLength * s.SideLength
}
