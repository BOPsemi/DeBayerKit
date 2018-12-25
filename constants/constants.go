package constants

/*
Definition of all constants
*/

/*
ImageFileFormat :image format constants
*/
type ImageFileFormat int

/*
PNG :PNG format
JPG :jpeg format
*/
const (
	PNG ImageFileFormat = iota
	JPG
)

/*
BayerArrayFormat :define bayer array format
*/
type BayerArrayFormat int

/*
RGGB :Red-Gr-Gb-Blue
GRBG :Gr-Red-Blue-Gb
BGGR :Blue-Gb-Gr-Red
*/
const (
	RGGB BayerArrayFormat = iota
	GRBG
	BGGR
)
