package converters

import (
	"DeBayerKit/constants"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

/*
RGBtoBayer :convert RGB stacked data to bayer
*/
type RGBtoBayer interface {
	ReArrangeToBayer() bool
}

type rgb struct {
	r uint32
	g uint32
	b uint32
}

/*
StackToBayerConverter :stacked image to bayer
*/
type StackToBayerConverter struct {
	RGBtoBayer

	img    image.Image                // original image
	format constants.BayerArrayFormat // bayer array constants
	bayer  [][]uint8                  // bayer data

	// image size
	imgSize struct {
		x int // width
		y int // height
	}
}

/*
NewStackToBayerConverter :initializer of object
*/
func NewStackToBayerConverter(img image.Image, format constants.BayerArrayFormat) *StackToBayerConverter {
	if img == nil {
		return nil
	}
	obj := new(StackToBayerConverter)

	obj.img = img
	obj.imgSize.x = img.Bounds().Size().X
	obj.imgSize.y = img.Bounds().Size().Y

	obj.format = format

	// debug
	//fmt.Println(obj.imgSize, obj.format)

	return obj
}

// read 2x2 image
func (s *StackToBayerConverter) readImageClustor(x, y int) []uint8 {

	origin := new(rgb)
	xshift := new(rgb)
	yshift := new(rgb)
	xyshift := new(rgb)

	origin.r, origin.g, origin.b, _ = s.img.At(x, y).RGBA()
	xshift.r, xshift.g, xshift.b, _ = s.img.At(x+1, y).RGBA()
	yshift.r, yshift.g, yshift.b, _ = s.img.At(x, y+1).RGBA()
	xyshift.r, xyshift.g, xyshift.b, _ = s.img.At(x+1, y+1).RGBA()

	switch s.format {
	case constants.RGGB:
		r := uint8(origin.r)
		gr := uint8(xshift.g)
		gb := uint8(yshift.g)
		b := uint8(xyshift.b)
		return []uint8{r, gr, gb, b}

	case constants.GRBG:
		gr := uint8(origin.g)
		r := uint8(xshift.r)
		b := uint8(yshift.b)
		gb := uint8(xyshift.b)
		return []uint8{gr, r, b, gb}

	case constants.BGGR:
		b := uint8(origin.b)
		gb := uint8(xshift.g)
		gr := uint8(yshift.g)
		r := uint8(xyshift.r)
		return []uint8{b, gb, gr, r}

	default:
		return nil
	}
}

// writeImageClustor
func (s *StackToBayerConverter) writeImageClustor() {
	canvas := image.NewRGBA(image.Rect(0, 0, s.imgSize.x, s.imgSize.y))

	heightMax := s.imgSize.y / 2
	for width := 0; width < s.imgSize.x/2; width++ {
		for height := 0; height < s.imgSize.y/2; height++ {

			i := width * 2
			j := height * 2

			bayerData := s.bayer[width*heightMax+height]
			r := bayerData[0]
			gr := bayerData[1]
			gb := bayerData[2]
			b := bayerData[3]

			canvas.Set(i, j, color.RGBA{r, 0, 0, 255})
			canvas.Set(i+1, j, color.RGBA{0, gr, 0, 255})
			canvas.Set(i, j+1, color.RGBA{0, gb, 0, 255})
			canvas.Set(i+1, j+1, color.RGBA{0, 0, b, 255})
		}
	}
}

/*
ReArrangeToBayer :implimentation of function
	in	:
	out	:
*/
func (s *StackToBayerConverter) ReArrangeToBayer() bool {

	var stocker [][]uint8

	for width := 0; width < s.imgSize.x/2; width++ {
		for height := 0; height < s.imgSize.y/2; height++ {
			i := width * 2
			j := height * 2
			pixel := s.readImageClustor(i, j)

			stocker = append(stocker, pixel)
		}
	}

	if stocker == nil {
		return false
	}

	s.bayer = stocker
	s.remakeImage()

	return true
}

func (s *StackToBayerConverter) remakeImage() bool {
	canvas := image.NewRGBA(image.Rect(0, 0, s.imgSize.x, s.imgSize.y))

	var lastNumber int
	heightMax := s.imgSize.y / 2
	for width := 0; width < s.imgSize.x/2; width++ {
		for height := 0; height < s.imgSize.y/2; height++ {
			lastNumber = width*heightMax + height

			i := width * 2
			j := height * 2

			bayerData := s.bayer[width*heightMax+height]
			r := bayerData[0]
			gr := bayerData[1]
			gb := bayerData[2]
			b := bayerData[3]

			canvas.Set(i, j, color.RGBA{r, 0, 0, 255})
			canvas.Set(i+1, j, color.RGBA{0, gr, 0, 255})
			canvas.Set(i, j+1, color.RGBA{0, gb, 0, 255})
			canvas.Set(i+1, j+1, color.RGBA{0, 0, b, 255})
		}
	}

	file, _ := os.Create("/Users/kazufumiwatanabe/go/src/DeBayerKit/data/test.png")
	defer file.Close()

	png.Encode(file, canvas)

	fmt.Println(lastNumber)

	return true
}
