package main

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

func isValid(s string) bool {
	left := []rune{'{', '[', '('}
	right := []rune{'}', ']', ')'}

	opened := make([]rune, 0, 0)

	for _, val := range s {
		if slices.Contains(left, val) {
			opened = append(opened, val)
		}
		if slices.Contains(right, val) {
			if len(opened) > 0 {
				idx := slices.Index(right, val)

				if left[idx] == opened[len(opened)-1] {
					opened = opened[:len(opened)-1]
				} else {
					return false
				}
			}

		}
	}

	if len(opened) > 0 {
		return false
	}

	return true
}

// Cases:
// fmt.Println(isValid("{[(())]}"))
// fmt.Println(isValid("{(())]}"))
// fmt.Println(isValid("{[(()]}"))
// fmt.Println(isValid("{[{[([()])]}{}()](({}))}"))

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
}

/*
TODO: different approach:

Create algoritht, that will find first `Closed` brace occurence and it will check if
previous one is the same type. Then it will `remove` them from arr, and go further.
If array will be empty, sequence was `valid`, in other case it was not
*/
