package converters

import (
	"DeBayer/constants"
	"image"
	"image/color"
)

/*
RGBtoBayer :convert RGB stacked data to bayer
*/
type RGBtoBayer interface {
	ReArrangeToBayer() bool
	DeBayerImage() image.Image
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

	deBayerImg image.Image // de-Bayer image
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

// read clustor
func (s *StackToBayerConverter) readRGGBCluster(origin, xshift, yshift, xyshift *rgb) []uint8 {
	r := uint8(origin.r)
	gr := uint8(xshift.g)
	gb := uint8(yshift.g)
	b := uint8(xyshift.b)

	return []uint8{r, gr, gb, b}
}

func (s *StackToBayerConverter) readGRBGCluster(origin, xshift, yshift, xyshift *rgb) []uint8 {
	gr := uint8(origin.g)
	r := uint8(xshift.r)
	b := uint8(yshift.b)
	gb := uint8(xyshift.b)
	return []uint8{gr, r, b, gb}
}

func (s *StackToBayerConverter) readBGGRCluster(origin, xshift, yshift, xyshift *rgb) []uint8 {
	b := uint8(origin.b)
	gb := uint8(xshift.g)
	gr := uint8(yshift.g)
	r := uint8(xyshift.r)
	return []uint8{b, gb, gr, r}
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
		return s.readRGGBCluster(origin, xshift, yshift, xyshift)
	case constants.GRBG:
		return s.readGRBGCluster(origin, xshift, yshift, xyshift)
	case constants.BGGR:
		return s.readBGGRCluster(origin, xshift, yshift, xyshift)
	default:
		return nil
	}
}

/*
ReArrangeToBayer :implimentation of function
	in	:
	out	:
*/
func (s *StackToBayerConverter) ReArrangeToBayer() bool {

	// make bayer data from original image
	var stocker [][]uint8
	for width := 0; width < s.imgSize.x/2; width++ {
		for height := 0; height < s.imgSize.y/2; height++ {
			i := width * 2
			j := height * 2
			pixel := s.readImageClustor(i, j)

			stocker = append(stocker, pixel)
		}
	}

	// check data set
	if stocker == nil {
		return false
	}

	// upload bayer data
	s.bayer = stocker

	// make de-Bayer image
	deBayerImg := s.makeBayerImage()
	if deBayerImg == nil {
		return false
	}
	s.deBayerImg = deBayerImg

	return true
}

// canvas writers
// For RGGB
func (s *StackToBayerConverter) rggbCanvasWriter() *image.RGBA {
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

	return canvas
}

// For GRBG
func (s *StackToBayerConverter) grbgCanvasWriter() *image.RGBA {
	canvas := image.NewRGBA(image.Rect(0, 0, s.imgSize.x, s.imgSize.y))

	heightMax := s.imgSize.y / 2
	for width := 0; width < s.imgSize.x/2; width++ {
		for height := 0; height < s.imgSize.y/2; height++ {

			i := width * 2
			j := height * 2

			bayerData := s.bayer[width*heightMax+height]
			gr := bayerData[0]
			r := bayerData[1]
			b := bayerData[2]
			gb := bayerData[3]

			canvas.Set(i, j, color.RGBA{0, gr, 0, 255})
			canvas.Set(i+1, j, color.RGBA{r, 0, 0, 255})
			canvas.Set(i, j+1, color.RGBA{0, 0, b, 255})
			canvas.Set(i+1, j+1, color.RGBA{0, gb, 0, 255})
		}
	}

	return canvas
}

// For BGGR
func (s *StackToBayerConverter) bggrCanvasWriter() *image.RGBA {
	canvas := image.NewRGBA(image.Rect(0, 0, s.imgSize.x, s.imgSize.y))

	heightMax := s.imgSize.y / 2
	for width := 0; width < s.imgSize.x/2; width++ {
		for height := 0; height < s.imgSize.y/2; height++ {

			i := width * 2
			j := height * 2

			bayerData := s.bayer[width*heightMax+height]
			b := bayerData[0]
			gb := bayerData[1]
			gr := bayerData[2]
			r := bayerData[3]

			canvas.Set(i, j, color.RGBA{0, 0, b, 255})
			canvas.Set(i+1, j, color.RGBA{0, gb, 0, 255})
			canvas.Set(i, j+1, color.RGBA{0, gr, 0, 255})
			canvas.Set(i+1, j+1, color.RGBA{r, 0, 0, 255})
		}
	}

	return canvas
}

func (s *StackToBayerConverter) makeBayerImage() *image.RGBA {
	switch s.format {
	case constants.RGGB:
		return s.rggbCanvasWriter()
	case constants.BGGR:
		return s.bggrCanvasWriter()
	case constants.GRBG:
		return s.grbgCanvasWriter()
	default:
		return nil
	}
}

/*
DeBayerImage :getter
	in	:
	out	:image.Image
*/
func (s *StackToBayerConverter) DeBayerImage() image.Image {
	return s.deBayerImg
}

/*
func (s *StackToBayerConverter) remakeImage() bool {
	canvas := s.rggbCanvasWriter()

	file, _ := os.Create("/Users/kazufumiwatanabe/go/src/DeBayer/data/out.png")
	defer file.Close()

	png.Encode(file, canvas)

	return true
}
*/
