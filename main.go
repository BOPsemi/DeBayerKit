package main

import (
	"DeBayer/constants"
	"DeBayer/converters"
	"flag"
)

func main() {
	var (
		inPath  = flag.String("inpath", "./data", "input file string path")
		outPath = flag.String("outpath", "./out", "output file string path")
		bayer   = flag.String("bayer", "RGGB", "Bayer pattern option")
	)
	// command line parse
	flag.Parse()

	converter := converters.NewBayerImage()
	if converter.SetReadWriteFileFolder(*inPath, *outPath) {
		switch *bayer {
		case "RGGB":
			converter.GenerateBayerImage(constants.RGGB)
		case "BGGR":
			converter.GenerateBayerImage(constants.BGGR)
		case "GRBG":
			converter.GenerateBayerImage(constants.GRBG)
		default:
			return
		}
	}
}
