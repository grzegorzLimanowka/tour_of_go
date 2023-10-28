package basics

import (
	"fmt"
	"math"
	"strings"
)

// In go you can define methods on types, via the special 'receiver' arg

type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale2(f float64) *Vertex2 {
	v.X *= f
	v.Y *= f

	v2 := Vertex2{v.X, v.Y}

	v3 := &v2

	return v3
}

// func (v *Vertex2) Scale(f float64) *Vertex2 {

// }

func (v *Vertex2) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main_vertex() {
	vertex2 := Vertex2{3, 4}
	fmt.Println(vertex2)

	fmt.Println(vertex2.Abs())

	vertex2.Scale(2)
	fmt.Println(vertex2)

	fmt.Println(vertex2.Scale2(3))
	fmt.Println(vertex2)
}

type MyFloat float64

func Abs() float64 {

	return 22.0
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func (v Vertex2) Sum() int {
	return int(v.X + v.Y)
}

func WordCount(s string) map[string]int {
	words := make(map[string]int)

	split := strings.Split(s, " ")

	for i := range split {
		_, v := words[split[i]]

		if v == false {
			words[split[i]] = 1
		} else {
			words[split[i]] += 1
		}

	}

	return words
}
