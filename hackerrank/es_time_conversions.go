package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 07:05:45PM -> 19:05:45
type TimeAmerican struct {
	H  [2]byte
	M  [2]byte
	S  [2]byte
	Pm bool // false - AM, true - PM
}

func timeConversion(s string) string {
	parts := strings.Split(s, ":")
	var output = []int{0, 0, 0}

	H, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])
	S, _ := strconv.Atoi(parts[2][:2])

	output[0] = H
	output[1] = M
	output[2] = S

	last := parts[2]

	if last[2] == 'A' && last[3] == 'M' {
		if H == 12 {
			output[0] = H - 12
		}
	}

	if last[2] == 'P' && last[3] == 'M' {
		if H >= 1 && H <= 11 {
			output[0] = H + 12
		}
	}

	return fmt.Sprintf("%02d:%02d:%02d", output[0], output[1], output[2])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
