package main

import (
	"fmt"
	"image"
	"image/color"
)

type Image struct {
	posX, posY, width, height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {

	return image.Rect(i.posX, i.posY, i.width, i.height) // FIXME
}

func (i Image) At(x, y int) color.Color {
	r := uint8(x * y % 255)
	g := uint8(x / y % 255)
	b := uint8(x*y - y%255)
	a := uint8((x * y) / 2 * 6 % 255)

	return color.RGBA{r, g, b, a}
}

func main() {
	m := Image{10, 10, 33, 35}

	fmt.Println(m.ColorModel())
	fmt.Println(m.Bounds())
	fmt.Println(m.At(100, 100))

	// pic.ShowImage(m)
}
