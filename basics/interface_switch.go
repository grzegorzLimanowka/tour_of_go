package basics

import "fmt"

func do(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v \n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long \n", v, len(v))
	default:
		fmt.Printf("I dont know about type %T", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
