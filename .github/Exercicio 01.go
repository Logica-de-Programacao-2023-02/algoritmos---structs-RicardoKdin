package main

import (
	"fmt"
	"math"
)

type circulo struct {
	raio float64
}

func main() {
	c := circulo{
		raio: 1.2,
	}
	fmt.Print(areacirculo(c))

}

func areacirculo(c circulo) float64 {
	area := math.Pi * c.raio * c.raio
	return area
}
