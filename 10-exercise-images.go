package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
	colors        uint8
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img *Image) At(x, y int) color.Color {
	return color.RGBA{img.colors + uint8(x), img.colors + uint8(y), 255, 255}
}

func main() {
	m := Image{100, 100, 200}
	pic.ShowImage(&m)
}
