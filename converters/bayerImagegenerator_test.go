package converters

import (
	"DeBayer/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	INPATH  = "/Users/kazufumiwatanabe/go/src/DeBayer/data"
	OUTPATH = "/Users/kazufumiwatanabe/go/src/DeBayer/out"
)

func TestNewBayerImage(t *testing.T) {
	obj := NewBayerImage()

	assert.NotNil(t, obj)
}

func TestSetReadWriteFileFolder(t *testing.T) {
	obj := NewBayerImage()

	status := obj.SetReadWriteFileFolder(INPATH, OUTPATH)
	assert.True(t, status)

	/*
		status = obj.SetOutImageSize(640, 0)
		assert.True(t, status)
	*/

	status = obj.GenerateBayerImage(constants.RGGB)
	assert.True(t, status)
}
