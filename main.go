package main

import "fmt"

type S struct {
	a, b, c string
}

func main() {
	x := interface{}(&S{"a", "b", "c"})
	y := interface{}(&S{"a", "b", "c"})

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v", y, y)
}
