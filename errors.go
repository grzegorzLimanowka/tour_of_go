package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(float64(x))
	}

	z := 1.0

	for i := 0; i < 10; i++ {
		prev_z := z
		z -= (z*z - x) / (2 * z)

		if z == prev_z {
			break
		}

	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2)) // should return err?
}
