package converters

import (
	"DeBayerKit/constants"
	"DeBayerKit/util"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSignalPicker(t *testing.T) {
	obj := NewSignalPicker()
	assert.NotNil(t, obj)
}

func TestPickRGBSignal(t *testing.T) {
	reader := util.NewIOReader()
	img := reader.ReadImageFile(SAMPLEIMAGEPATH, constants.JPG)

	// check image reading
	assert.NotNil(t, img)

	// picker
	picker := NewSignalPicker()
	r, g, b := picker.PickRGBSignal(img, 0, 10)

	assert.NotEqual(t, 255, r)
	assert.NotEqual(t, 255, g)
	assert.NotEqual(t, 255, b)

	fmt.Println(r, g, b)
}
