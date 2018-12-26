package util

import (
	"DeBayer/constants"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

/*
Reader :interface of IO Reader
*/
type Reader interface {
	ReadImageFile(path string, format constants.ImageFileFormat) image.Image
	FilesInFolder(path string) []string
}

/*
Writer :interface of IO Writer
*/
type Writer interface {
	WrireImageFile(path string, img image.Image, format constants.ImageFileFormat) bool
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

// create file
func (i *IOTool) create(path string) *os.File {
	file, err := os.Create(path)
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
IOWriter :writer object
*/
type IOWriter struct {
	IOTool
	Writer // interface
}

/*
NewIOReader :initializer of IOReader
*/
func NewIOReader() *IOReader {
	return new(IOReader)
}

/*
NewIOWriter :initializer of IOWriter
*/
func NewIOWriter() *IOWriter {
	return new(IOWriter)
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

	// open file
	file := i.open(path)
	defer file.Close()

	switch format {
	case constants.JPG:
		// read jpg file
		img, err := jpeg.Decode(file)
		if err != nil {
			i.errorHandler(err)
			return nil
		}
		return img

	case constants.PNG:
		// read png file
		img, err := png.Decode(file)
		if err != nil {
			i.errorHandler(err)
			return nil
		}
		return img

	default:
		return nil
	}
}

/*
FilesInFolder :get file names in folder
	in	:path string
	out	:[]string
*/
func (i *IOReader) FilesInFolder(path string) []string {
	if path == "" {
		return nil
	}

	// read dir info
	files, err := ioutil.ReadDir(path)
	if err != nil {
		i.errorHandler(err)
		return nil
	}

	// get file names
	var list []string
	for _, file := range files {
		list = append(list, filepath.Join(path, file.Name()))
	}

	return list
}

/*
WrireImageFile :implimentaion of image writer
	in	:path string, img image.Image, format constants.ImageFileFormat
	out	:bool
*/
func (i *IOWriter) WrireImageFile(path string, img image.Image, format constants.ImageFileFormat) bool {
	if path == "" {
		return false
	}

	if img == nil {
		return false
	}

	// create file
	file := i.create(path)
	defer file.Close()

	// crate file
	switch format {
	case constants.PNG:
		err := png.Encode(file, img)
		if err != nil {
			i.errorHandler(err)
			return false
		}
		return true
	case constants.JPG:
		err := jpeg.Encode(file, img, nil)
		if err != nil {
			i.errorHandler(err)
			return false
		}
		return true
	default:
		return false
	}
}
