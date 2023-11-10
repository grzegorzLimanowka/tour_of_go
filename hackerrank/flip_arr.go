package main

import "fmt"

/*
 * Complete the 'flippingMatrix' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY matrix as parameter.
 */

func sum(input []int32) int32 {
	var sum int32 = 0
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}

	return sum
}

func flippingMatrix(matrix [][]int32) int32 {
	n := len(matrix) / 2

	fmt.Println(n)

	for i := 0; i < len(matrix); i++ {
		left := sum(matrix[i][:n])
		right := sum(matrix[i][n:])

		fmt.Printf("[i%v]: LEFT: %v RIGHT: %v\n", i, left, right)

		if right > left {
			fmt.Printf("B SWAP [%v] = %v\n", i, matrix[i])

			for j := 0; j < len(matrix[i])/2; j++ {
				// matrix[i][j], matrix[i][(len(matrix[j])-1)-1] = matrix[i][(len(matrix[j])-1)-1], matrix[i][j]
				fmt.Println("SWP", matrix[i][j], matrix[i][len(matrix[i])-1])

				matrix[i][j], matrix[i][(len(matrix[i])-1)-j] = matrix[i][(len(matrix[i])-1)-j], matrix[i][j]
			}

			fmt.Printf("A SWAP [%v] = %v\n --- \n", i, matrix[i])
		}
	}

	fmt.Println("HALF: Matrix", matrix)

	for i := 0; i < len(matrix); i++ {
		var left int32 = 0
		for j := 0; j < (len(matrix) / 2); j++ {
			left += matrix[j][i]
		}

		var right int32 = 0
		for j := len(matrix) / 2; j < len(matrix); j++ {
			right += matrix[j][i]
		}

		fmt.Printf("[i%v]: LEFT: %v RIGHT: %v\n", i, left, right)

		if right > left {
			fmt.Printf("B SWAP [%v] = [", i)

			for j := 0; j < len(matrix); j++ {
				fmt.Print(matrix[j][i])
			}
			fmt.Printf("]\n")

			for j := 0; j < len(matrix[i])/2; j++ {
				matrix[j][i], matrix[(len(matrix[i])-j)-1][i] = matrix[(len(matrix[i])-j)-1][i], matrix[j][i]
			}

			fmt.Printf("A SWAP [%v] = [", i)
			for j := 0; j < len(matrix); j++ {
				fmt.Print(matrix[j][i])
			}
			fmt.Printf("] \n ---- \n")
		}

	}

	fmt.Println("END: Matrix", matrix)

	var sum int32 = 0

	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix)/2; j++ {
			sum += matrix[i][j]
		}
	}

	return sum
}

func main() {
	ex := [][]int32{
		{12, 2, 3, 19},
		{6, 25, 6, 9},
		{5, 8, 1, 3},
		{2, 8, 14, 8},
	}

	out := flippingMatrix(ex)
	fmt.Println(out)

	ex2 := [][]int32{
		{112, 42, 83, 119},
		{56, 125, 56, 49},
		{15, 78, 101, 43},
		{62, 98, 114, 108},
	}
	out2 := flippingMatrix(ex2)
	fmt.Println(out2)

	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	// qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// checkError(err)
	// q := int32(qTemp)

	// for qItr := 0; qItr < int(q); qItr++ {
	// 	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// 	checkError(err)
	// 	n := int32(nTemp)

	// 	var matrix [][]int32
	// 	for i := 0; i < 2*int(n); i++ {
	// 		matrixRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

	// 		var matrixRow []int32
	// 		for _, matrixRowItem := range matrixRowTemp {
	// 			matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
	// 			checkError(err)
	// 			matrixItem := int32(matrixItemTemp)
	// 			matrixRow = append(matrixRow, matrixItem)
	// 		}

	// 		if len(matrixRow) != 2*int(n) {
	// 			panic("Bad input")
	// 		}

	// 		matrix = append(matrix, matrixRow)
	// 	}

	// 	result := flippingMatrix(matrix)

	// 	fmt.Fprintf(writer, "%d\n", result)
	// }

	// writer.Flush()
}

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// for j := 0; j < len(matrix); j++ {
// if sum(matrix[i][:n+1]) < sum(matrix[i][n+1:]) {
// 	fmt.Printf("B SWAP [%v][%v] = %v\n", i, j, matrix[i])

// 	for i := 0; i < len(matrix[i])/2; i++ {
// 		matrix[i][i], matrix[i][(len(matrix[i])-1)-1] = matrix[i][(len(matrix[i])-1)-1], matrix[i][i]
// 	}

// 	fmt.Printf("A SWAP [%v][%v] = %v\n --- \n", i, j, matrix[i])
// }
// }
