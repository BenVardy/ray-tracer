package core

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// NewImage creates a blank canvas
func NewImage(width, height int) *Image {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	return &Image{width, height, img}
}

// Image wraps around the image.RGBA struct to make getting and setting pixels easier
type Image struct {
	Width  int
	Height int
	Img    *image.RGBA
}

// SetPixel sets the pixel at (x, y) to the color c
func (i *Image) SetPixel(x, y int, c color.RGBA) {
	i.Img.SetRGBA(x, y, c)
}

// GetPixel returns the color at (x, y)
func (i *Image) GetPixel(x, y int) color.RGBA {
	return i.Img.RGBAAt(x, y)
}

// PrintToFile outputs the contents oof the Image to the file with name fname
func (i *Image) PrintToFile(fname string) error {
	f, err := os.Create(fname)
	defer f.Close()

	if err != nil {
		return err
	}

	png.Encode(f, i.Img)
	return nil
}
