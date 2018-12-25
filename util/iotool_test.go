package util

import (
	"DeBayerKit/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	SAMPLEIMAGEPATH = "/Users/kazufumiwatanabe/go/src/DeBayerKit/data/IMG_0023.jpeg"
)

func TestNewIOReader(t *testing.T) {
	obj := NewIOReader()
	assert.NotNil(t, obj)
}

func TestReadImageFile(t *testing.T) {
	obj := NewIOReader()
	img := obj.ReadImageFile(SAMPLEIMAGEPATH, constants.JPG)

	assert.NotNil(t, img)
}
