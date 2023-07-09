package shape

import "fmt"

type Shape interface {
	GetArea() float64
}

func PrintArea(s Shape) {
	fmt.Println(s.GetArea())
}
