package util

import (
	"DeBayer/constants"
	"image"
)

/*
FormatChange :interface of file format change
*/
type FormatChange interface {
	JPGtoPNG(inPath, outPath string) bool
	PNGtoJPG(inPath, outPath string) bool
}

/*
ImgFile :outline of image file converter
*/
type ImgFile struct {
	reader *IOReader // reader
	writer *IOWriter // writer
}

func (i *ImgFile) open(path string, format constants.ImageFileFormat) image.Image {
	if path == "" {
		return nil
	}

	// initializer reader object
	if i.reader == nil {
		i.reader = NewIOReader()
	}

	// open file
	return i.reader.ReadImageFile(path, format)
}

func (i *ImgFile) write(path string, img image.Image, format constants.ImageFileFormat) bool {
	if path == "" {
		return false
	}

	// initialize writer object
	if i.writer == nil {
		i.writer = NewIOWriter()
	}

	// write file
	return i.writer.WrireImageFile(path, img, format)
}

/*
ImgFormatChange :stuct of image format changer
*/
type ImgFormatChange struct {
	ImgFile      // parents
	FormatChange // interface
}

/*
NewImgFormatChange :initializer
*/
func NewImgFormatChange() *ImgFormatChange {
	return new(ImgFormatChange)
}

/*
JPGtoPNG :implimentation of interface
	in	:inPath, outPath string
	out	:bool
*/
func (i *ImgFormatChange) JPGtoPNG(inPath, outPath string) bool {
	if inPath == "" || outPath == "" {
		return false
	}

	// format change
	inImg := i.open(inPath, constants.JPG)
	if !i.write(outPath, inImg, constants.PNG) {
		return false
	}

	return true
}

/*
PNGtoJPG :implimentation of interface
	in	:inPath, outPath string
	out	:bool
*/
func (i *ImgFormatChange) PNGtoJPG(inPath, outPath string) bool {
	if inPath == "" || outPath == "" {
		return false
	}

	// format change
	inImg := i.open(inPath, constants.PNG)
	if !i.write(outPath, inImg, constants.JPG) {
		return false
	}

	return true
}
