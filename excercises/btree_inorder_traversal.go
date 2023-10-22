package excercises

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// InOrder BTree Traversal
func WalkBranch(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		WalkBranch(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		WalkBranch(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	WalkBranch(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) bool {
	fmt.Println("XXX")

	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// TODO: Find a way to better "sync" receiving values from 2 channels
	for x := range ch1 {
		y := <-ch2

		fmt.Println("x", x)
		fmt.Println("y", y)

		if x != y {
			return false
		}
	}

	return true
}

func main() {
	t1 := tree.New(10)
	fmt.Println(t1)

	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for x := range ch {
		fmt.Println(x)
	}

	fmt.Println(Same(tree.New(1), tree.New(1))) // should return true,
	fmt.Println(Same(tree.New(1), tree.New(2))) // should return false.
}
