package main

import "fmt"

type Triangulo struct {
	Base   float64
	Altura float64
}

func main() {
	tri := Triangulo{
		Base:   2.3,
		Altura: 3.2,
	}
	fmt.Print(areatriangulo(tri))
}

func areatriangulo(t Triangulo) float64 {
	area := t.Base * t.Altura / 2
	return area
}
