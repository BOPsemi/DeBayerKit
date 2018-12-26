package util

import (
	"image"

	"github.com/disintegration/imaging"
)

/*
Resize :interface of image size resizer
*/
type Resize interface {
	SetImageSize(width, height int) bool
	ResizeImageSize(img image.Image) image.Image
}

/*
ResizeImage :resize image struture
*/
type ResizeImage struct {
	Resize // interface

	width  int // target image width
	height int // target image height

}

/*
NewResizeImage :initializer of ResizeImage
*/
func NewResizeImage() *ResizeImage {
	obj := new(ResizeImage)

	return obj
}

/*
SetImageSize :set target image size
	in	:width, height int
	out	:bool
*/
func (r *ResizeImage) SetImageSize(width, height int) bool {
	if width <= 0 {
		return false
	}

	// set width and height
	r.width = width
	r.height = height

	return true
}

/*
ResizeImageSize :resize image size
	in	:img image.Image
	out	:image.Image
*/
func (r *ResizeImage) ResizeImageSize(img image.Image) image.Image {
	if img == nil {
		return nil
	}

	// initialize tool
	resizedImage := imaging.Resize(img, r.width, r.height, imaging.Lanczos)
	if resizedImage == nil {
		return nil
	}

	return resizedImage
}
