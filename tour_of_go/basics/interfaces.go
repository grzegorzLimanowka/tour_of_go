package basics

import (
	"fmt"
	"math"
)

// In go you can define methods on types, via the special 'receiver' arg

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt(2))
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// in the following line, v is a Vertex (not a *Vertex)
	// and does NOT implement Abser
	// a = v

	fmt.Println(a.Abs())

}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
