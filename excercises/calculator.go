package excercises

// Simple calculator, just to fiddle around with golangs interfaces

import "fmt"

type Calculator interface {
	Add(x float64) float64
	Substract(x float64) float64
	Multiply(x float64) float64
	Divide(x float64) float64
}

type SimpleCalculator struct {
	sum float64
}

func (c *SimpleCalculator) Add(x float64) float64 {
	c.sum += x

	return c.sum
}

func (c *SimpleCalculator) Substract(x float64) float64 {
	c.sum -= x

	return c.sum

}

func (c *SimpleCalculator) Multiply(x float64) float64 {
	c.sum *= x

	return c.sum

}

func (c *SimpleCalculator) Divide(x float64) float64 {
	if x == 0 {
		fmt.Println("Cannot divide by 0 !!!")

		return c.sum
		// panic("Calculator panicked")
	}

	c.sum /= x

	return c.sum

}

func main() {
	var calc Calculator = &SimpleCalculator{0}

	fmt.Println(calc.Add(44))
	fmt.Println(calc.Substract(22))
	fmt.Println(calc.Multiply(44))
	fmt.Println(calc.Divide(44))
}
