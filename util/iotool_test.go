package util

import (
	"DeBayer/constants"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	SAMPLEIMAGEPATH = "/Users/kazufumiwatanabe/go/src/DeBayer/data/IMG_0870.png"
	INIMG           = "/Users/kazufumiwatanabe/go/src/DeBayer/data/DSC01899.jpg"
	OUTIMG          = "/Users/kazufumiwatanabe/go/src/DeBayer/data/DSC01899.png"
)

func TestNewIOReader(t *testing.T) {
	obj := NewIOReader()
	assert.NotNil(t, obj)
}

func TestReadImageFile(t *testing.T) {
	obj := NewIOReader()
	img := obj.ReadImageFile(SAMPLEIMAGEPATH, constants.PNG)

	assert.NotNil(t, img)
}

func TestFilesInFolder(t *testing.T) {
	obj := NewIOReader()
	list := obj.FilesInFolder("/Users/kazufumiwatanabe/go/src/DeBayer/data")

	assert.NotNil(t, list)
	assert.Equal(t, 2, len(list))

	fmt.Println(list)
}

func TestImgFormatChange(t *testing.T) {
	obj := NewImgFormatChange()

	assert.NotNil(t, obj)

	status := obj.JPGtoPNG(INIMG, OUTIMG)
	assert.True(t, status)

}
