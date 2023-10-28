package basics

import "fmt"

const Pi = 3.14

const (
	// Create huge number by shifting a 1 bit left 10 places
	// In other words, the binary number that is 1 followed by 10 zeroes,
	// So it should be 0100 0000 0000, so 1024
	Small = 1 << 10

	Big = 1 << 100

	// Stift it right again 99 places, so we end up with 1<<1, so 2
	VerySmall = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func xx() (z int, y int) {
	z = 2
	y = 2
	return
}

func main333() {
	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(VerySmall))
	fmt.Println(needFloat(VerySmall))
}
