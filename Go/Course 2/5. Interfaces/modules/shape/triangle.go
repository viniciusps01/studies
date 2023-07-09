package shape

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) GetArea() float64 {
	return t.Base * t.Height * .5
}
