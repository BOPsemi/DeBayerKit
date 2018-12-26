package main

import (
	"DeBayer/util"
	"flag"
	"fmt"
)

func main() {
	var (
		path = flag.String("path", "./data", "string path")
	)
	// command line parse
	flag.Parse()
	fmt.Println(*path)

	// get all file info
	reader := util.NewIOReader()
	list := reader.FilesInFolder(*path)

	fmt.Println(list)
}
