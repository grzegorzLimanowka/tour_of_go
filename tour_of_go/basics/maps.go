package basics

import "fmt"

type Vertex struct {
	Lat, Long float64
}

type Vertex2 struct {
	Lattitude, Longitude float64
}

var m map[string]Vertex
var m2 map[int64]Vertex
var m3 map[int16]Vertex2

func maps() {
	m = make(map[string]Vertex)
	m["Bell labs"] = Vertex{
		40.6432, -42.1452,
	}

	m3 = make(map[int16]Vertex2)

	m3[23] = Vertex2{
		32.23, 234.24,
	}

	fmt.Println(m)

}

type Vertex3 struct {
	Lat, Long float64
}

func main() {

	var m2 = map[string]Vertex3{
		"Bell Labs": {
			32.42, -23.222,
		},
		"Google": {
			23.442, 2233.222,
		},
	}

	fmt.Println(m2)

	// --

	for k, v := range m {
		fmt.Printf("Key: %v, val: %v \n", k, v)
	}

	m := make(map[string]int)

	elem, ok := m["Answer"]
	fmt.Println(elem, ok)

	m["Answer"] = 42
	fmt.Println("The value", m["Answer"])

	elem, ok = m["Answer"]
	fmt.Println(elem, ok)

	m["Answer"] = 48
	fmt.Println("The value", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value", m["Answer"])
	fmt.Println("The value", m["Answer2"])

	elem, ok = m["Answer"]
	fmt.Println(elem, ok)

}
