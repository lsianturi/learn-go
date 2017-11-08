package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
	c    uint8
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) At(X, Y int) color.Color {
	return color.RGBA{i.c + uint8(X*Y+X*Y), i.c + uint8(X*Y+X*Y), 255, 255}

}

func main() {
	m := Image{200, 200, 128}
	pic.ShowImage(&m)
}
