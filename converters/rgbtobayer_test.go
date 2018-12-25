package converters

import (
	"DeBayerKit/constants"
	"DeBayerKit/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	SAMPLEIMAGEPATH = "/Users/kazufumiwatanabe/go/src/DeBayerKit/data/IMG_0870.png"
)

func TestNewNewStackToBayerConverter(t *testing.T) {
	reader := util.NewIOReader()
	img := reader.ReadImageFile(SAMPLEIMAGEPATH, constants.PNG)

	// check img read
	assert.NotNil(t, img)

	// check object creation
	obj := NewStackToBayerConverter(img, constants.RGGB)
	assert.NotNil(t, obj)

	// check function
	obj.ReArrangeToBayer()

}
