package main

import (
	"fmt"
	"strings"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	board := [][]string{
		[]string{".", ".", "."},
		[]string{".", ".", "."},
		[]string{".", ".", "."},
	}

	// Players take turns:
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s \n", strings.Join(board[i], ""))
	}

	for i, v := range pow {
		fmt.Printf("2**%d = %d \n", i, v)
	}

	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v \n", len(s), cap(s), s)
}
