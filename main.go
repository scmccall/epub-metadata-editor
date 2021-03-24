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

	// Get .epub file
	// Create temp directory to store unzipped files into
	// Unzip .epub contents into temp directory
	// Edit metadata
	// zip files from temp directory into new .epub file

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

func CreateTempDir() {
	tempDir, err := ioutil.TempDir("", "go-epub")
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			panic(fmt.Sprintf("Error removing temp directory: %s", err))
			return e go
		}
	}()
	if err != nil {
		panic(fmt.Sprintf("Error creating temp directory: %s", err))
		return err
	}
	return tempDir
}
