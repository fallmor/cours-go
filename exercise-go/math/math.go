package main

import (
	"fmt"
	"math"
)

type cercle struct {
	rayon float64
}
type rectangle struct {
	longueur float64
	largeur  float64
}

func (c cercle) surface() float64 {
	x := math.Pi * 2 * c.rayon
	return x
}
func (r rectangle) surface() float64 {
	x := r.longueur * r.largeur
	return x
}

type forme interface {
	surface() (string float64)
}

func foo(f forme) {
	fmt.Println(f.surface())
}

func main() {
	p1 := cercle{
		rayon: 12.3,
	}
	p2 := rectangle{
		longueur: 10,
		largeur:  10,
	}
	foo(p1)
	foo(p2)
}
