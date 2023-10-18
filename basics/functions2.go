package basics

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func compute2(fn func(float64) float64) float64 {
	return fn(2)
}

func main12() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	my := func(x float64) float64 {
		return 1 / x
	}

	fmt.Println(my(44))
	fmt.Println(compute2(my))

}
