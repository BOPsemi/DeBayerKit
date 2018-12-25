package util

import (
	"DeBayerKit/constants"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

/*
Reader :interface of IO Reader
*/
type Reader interface {
	ReadImageFile(path string, format constants.ImageFileFormat) image.Image
}

/*
IOTool : top IOTool structure
*/
type IOTool struct {
}

func (i *IOTool) errorHandler(err error) {
	log.Println(err.Error())
}

// open file
func (i *IOTool) open(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil
	}
	return file
}

/*
IOReader :reader object
*/
type IOReader struct {
	IOTool
	Reader // interface
}

/*
NewIOReader :initializer of IOReader
*/
func NewIOReader() *IOReader {
	return new(IOReader)
}

/*
ReadImageFile // read image file
	in	:path string, format constants.ImageFileFormat
	out	:image.Image
*/
func (i *IOReader) ReadImageFile(path string, format constants.ImageFileFormat) image.Image {
	if path == "" {
		return nil
	}

	switch format {
	/*
		JPG file format case
	*/
	case constants.JPG:

		// open file
		file := i.open(path)
		defer file.Close()

		// read jpg file
		img, err := jpeg.Decode(file)
		if err != nil {
			i.errorHandler(err)
		}
		return img

	/*
		PNG file format case
	*/
	case constants.PNG:

		// open file
		file := i.open(path)
		defer file.Close()

		// read png file
		img, err := png.Decode(file)
		if err != nil {
			i.errorHandler(err)
		}
		return img

	default:
		return nil
	}
}
