package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FileLocations struct {
	src  string
	ext  string
	dest string
}

func main() {

	Write()

}

func Write() error {

	// Get file name

	// While running "go run main.go zip.go write.go -- [name_of_file].epub" use Args[2]
	// For final production, use Args[1]
	fileName := os.Args[2]
	// Remove extension
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

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
	unzipEpub(fileName, tempDir)
	if err != nil {
		return err
	}

	// Get absolute filepath for temp directory
	tempDirAbsPath, err := filepath.Abs(tempDir)
	if err != nil {
		return err
	}

	// // Edit metadata
	editMetadata(tempDirAbsPath)

	// zip files from temp directory into new .epub file
	writeEpub(tempDir, fileName+".epub")

	return nil

}
