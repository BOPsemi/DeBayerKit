package util

import (
	"DeBayer/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	SIZINGIMGPATH  = "/Users/kazufumiwatanabe/go/src/DeBayer/data/DSC01899.png"
	RESIZEDIMGPATH = "/Users/kazufumiwatanabe/go/src/DeBayer/data/DSC01899_resized.png"
)

func TestNewResizeImage(t *testing.T) {
	obj := NewResizeImage()
	assert.NotNil(t, obj)
}

func TestResizeImage(t *testing.T) {
	obj := NewResizeImage()

	status := obj.SetImageSize(0, 0)
	assert.False(t, status)

	status = obj.SetImageSize(400, 0)
	assert.True(t, status)

	reader := NewIOReader()
	img := reader.ReadImageFile(SIZINGIMGPATH, constants.PNG)

	assert.NotNil(t, img)

	resizedImg := obj.ResizeImageSize(img)
	assert.NotNil(t, resizedImg)

	writer := NewIOWriter()
	status = writer.WrireImageFile(RESIZEDIMGPATH, resizedImg, constants.PNG)
	assert.True(t, status)
}
