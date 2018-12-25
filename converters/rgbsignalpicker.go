package converters

import "image"

/*
RGBSignalPicker :RGB signal picker
*/
type RGBSignalPicker interface {
	PickRGBSignal(img image.Image, x, y int) (r, g, b uint8)
}

/*
SignalPicker :sinal picker model
*/
type SignalPicker struct {
	RGBSignalPicker // interface
}

/*
NewSignalPicker :initializer
*/
func NewSignalPicker() *SignalPicker {
	obj := new(SignalPicker)
	return obj
}

/*
PickRGBSignal :interface impliment
	in	:img image.Image, x, y int
	out	:r, g, b uint8
*/
func (s *SignalPicker) PickRGBSignal(img image.Image, x, y int) (r, g, b uint8) {
	width := img.Bounds().Size().X
	height := img.Bounds().Size().Y

	if x > width || y > height {
		return 255, 255, 255
	}

	red, green, blue, _ := img.At(x, y).RGBA()
	return uint8(red), uint8(green), uint8(blue)

}
