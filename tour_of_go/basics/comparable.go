package basics

import "fmt"

func Index[T comparable](data []T, x T) int {
	for i, v := range data {
		if v == x {
			return i
		}
	}

	return -1
}

func main4() {
	// Index works on a slice if ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	//index also works on a slice of strings"
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

}
