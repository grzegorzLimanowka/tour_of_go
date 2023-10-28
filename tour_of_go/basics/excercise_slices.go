package basics

func Pic2(dx, dy int) [][]uint8 {

	arr := make([][]uint8, dx)

	for x := 0; x < dx; x++ {
		line := make([]uint8, dy)
		arr[x] = line

		for y := 0; y < dy; y++ {
			arr[x][y] = uint8((x * y) / 2)
		}

	}

	return arr
}
