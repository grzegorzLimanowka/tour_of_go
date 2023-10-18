package basics

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}    //
	v2 = Vertex{X: 1}    // y implicit 0
	v3 = Vertex{Y: 44}   // x implicit 0
	v4 = Vertex{}        // implicit x, y = 0
	p  = &Vertex{1, 444} // has type *Vertex
)

func main2() {
	fmt.Println(v1, p, v2, v3, v4)

	// Arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[0:6]
	fmt.Println(s)

	var c []int = primes[1:5]
	fmt.Println(c)
	c[0] = 44

	var d []int = primes[0:6]
	fmt.Println(d)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	b := names[1:4]

	b[0] = "xxx"

	fmt.Println(a, b)
	fmt.Println(names)

}

type Marian struct {
	x int
	y int
}

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)

	t := []struct {
		i float64
		b bool
	}{
		{float64(q[0]), r[0]},
		{float64(q[1]), r[1]},
		{float64(q[2]), r[2]},
		{float64(q[3]), r[3]},
	}

	fmt.Println(t)

}
