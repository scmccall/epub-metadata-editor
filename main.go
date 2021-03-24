package main

import (
// "log"
// "strings"
)

type FileLocations struct {
	src  string
	ext  string
	dest string
}

func main() {

	fileName := "songbird"
	temp := FileLocations{
		src:  fileName,
		ext:  ".epub",
		dest: fileName,
	}

	// Unzip the zip/epub file
	UnzipHelper(temp.src, temp.ext, temp.dest)

	// Zip the modified directory
	ZipHelper(temp.src, temp.ext)
}

func UnzipHelper(src string, ext string, dest string) error {
	_, err := Unzip(src, ext, dest)
	if err != nil {
		return err
	}
	return nil
}

func ZipHelper(src string, ext string) error {
	err := Zip(src, ext)
	if err != nil {
		return err
	}
	return nil
}
