package converters

import (
	"DeBayer/constants"
	"DeBayer/util"
	"image"
	"strings"
)

/*
Generator :image converter interface
*/
type Generator interface {
	SetReadWriteFileFolder(inpath, outpath string) bool
	SetOutImageSize(width, height int) bool
	GenerateBayerImage(bayerArray constants.BayerArrayFormat) bool
}

/*
BayerImage : bayer image genertor
*/
type BayerImage struct {
	Generator // interface

	inpath   string   // input path
	outpath  string   // output path
	list     []string // file list in input path
	saveList []string // renamed file list, extention was changed (PNG)

	width  int // output image file width
	height int // output image file height

}

/*
NewBayerImage : initializer
*/
func NewBayerImage() *BayerImage {
	obj := new(BayerImage)

	// set default value
	obj.width = -999

	return obj
}

/*
SetReadWriteFileFolder :set read and write file folder path
	in	:inpath, outpath string
	out	:bool
*/
func (b *BayerImage) SetReadWriteFileFolder(inpath, outpath string) bool {
	if inpath == "" || outpath == "" {
		return false
	}

	// set in and out path
	b.inpath = inpath
	b.outpath = outpath

	// read files in inpath
	reader := util.NewIOReader()
	b.list = reader.FilesInFolder(b.inpath)

	// save path setting
	var savePath string
	if !strings.HasSuffix(outpath, "/") {
		savePath = outpath + "/"
	} else {
		savePath = outpath
	}

	for _, name := range b.list {

		// extract file name
		fullNames := strings.Split(name, "/")
		fileName := fullNames[len(fullNames)-1]

		saveFilePath := savePath + fileName
		b.saveList = append(b.saveList, saveFilePath)
	}

	return true
}

/*
SetOutImageSize :set image size
	in	:width, height int
	out	:bool
*/
func (b *BayerImage) SetOutImageSize(width, height int) bool {
	if width <= 0 {
		return false
	}

	b.width = width
	b.height = height

	return true
}

/*
GenerateBayerImage :generte bayer image
	in	:
	out	:bool
*/
func (b *BayerImage) GenerateBayerImage(bayerArray constants.BayerArrayFormat) bool {
	// initialize reader and writer
	reader := util.NewIOReader()
	writer := util.NewIOWriter()

	// read image and prepare files for bayer
	for index, path := range b.list {
		var img image.Image

		// read image
		img = reader.ReadImageFile(path, constants.PNG)

		// generate bayer image
		converter := NewStackToBayerConverter(img, bayerArray)
		if converter.ReArrangeToBayer() {
			bayerImg := converter.DeBayerImage()
			writer.WrireImageFile(b.saveList[index], bayerImg, constants.PNG)
		}
	}

	return true
}
