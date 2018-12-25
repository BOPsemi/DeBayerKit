package util

import (
	"DeBayer/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	SAMPLEIMAGEPATH = "/Users/kazufumiwatanabe/go/src/DeBayer/data/IMG_0870.png"
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
