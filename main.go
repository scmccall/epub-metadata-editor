package main

import (
	// "log"
	// "strings"
	"fmt"
	"io/ioutil"
	"os"
)

type FileLocations struct {
	src  string
	ext  string
	dest string
}

func main() {

	Write()

}

// 	// Get .epub file
// 	// Create temp directory to store unzipped files into
// 	// Unzip .epub contents into temp directory
// 	// Edit metadata
// 	// zip files from temp directory into new .epub file

// 	fileName := "songbird"
// 	temp := FileLocations{
// 		src:  fileName,
// 		ext:  ".epub",
// 		dest: fileName,
// 	}

// 	// Create temp directory
// 	tempDir, err := ioutil.TempDir("", "go-epub")
// 	defer func() {
// 		if err := os.RemoveAll(tempDir); err != nil {
// 			panic(fmt.Sprintf("Error removing temp directory: %s", err))
// 		}
// 	}()
// 	if err != nil {
// 		panic(fmt.Sprintf("Error creating temp directory: %s", err))
// 	}

// 	// Unzip the zip/epub file
// 	UnzipHelper(temp.src, temp.ext, tempDir)

// 	// Write epub file
// 	writeEpub(tempDir, temp.src+temp.ext)

// 	// // Zip the modified directory
// 	// ZipHelper(temp.src, temp.ext)
// }

func Write() error {

	// Get .epub file
	fileName := "songbird"
	temp := FileLocations{
		src:  fileName,
		ext:  ".epub",
		dest: fileName,
	}

	// Create temp directory to store unzipped files into
	tempDir, err := ioutil.TempDir("", "go-epub")
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			panic(fmt.Sprintf("Error removing temp directory: %s", err))
		}
	}()
	if err != nil {
		panic(fmt.Sprintf("Error creating temp directory: %s", err))
	}

	// Unzip .epub contents into temp directory
	Unzip(temp.src, temp.ext, tempDir)
	if err != nil {
		return err
	}

	// Edit metadata

	// zip files from temp directory into new .epub file
	writeEpub(tempDir, temp.src+temp.ext)

	return nil

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
