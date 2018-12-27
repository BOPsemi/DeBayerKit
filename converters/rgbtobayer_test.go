package converters

import (
	"DeBayer/constants"
	"DeBayer/util"
	"image/png"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	SAMPLEIMAGEPATH = "/Users/kazufumiwatanabe/go/src/DeBayer/out/IMG_0028.png"
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

	// check deBayerImg
	deBayerImg := obj.DeBayerImage()
	assert.NotNil(t, deBayerImg)

	file, _ := os.Create("/Users/kazufumiwatanabe/go/src/DeBayer/data/out.png")
	defer file.Close()

	png.Encode(file, deBayerImg)

}
