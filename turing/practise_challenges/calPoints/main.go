package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calPoints(ops []string) int {
	var res int = 0
	record := make([]int, 0, 0)

	for i := 0; i < len(ops); i++ {
		point, err := strconv.Atoi(ops[i])

		if err == nil {
			record = append(record, point)
		} else {
			if ops[i] == "C" {
				record = record[:len(record)-1] // removed last
			}
			if ops[i] == "D" {
				new := record[len(record)-1] * 2 // we assume the exist

				record = append(record, new)
			}
			if ops[i] == "+" {
				new := record[len(record)-1] + record[len(record)-2] // we assume the exist

				record = append(record, new)
			}
		}
	}

	for i := 0; i < len(record); i++ {
		res += record[i]
	}

	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')
	rawInput = strings.Replace(rawInput, "\n", "", -1)

	ops := strings.Split(rawInput, " ")

	fmt.Println(calPoints(ops))
}
