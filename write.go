package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func writeEpub(tempDir string, destFilePath string) error {

	var mimetypeFilename = "mimetype"

	f, err := os.Create(destFilePath)
	if err != nil {
		panic(fmt.Sprintf("Error creating EPUB at %q: %+v", destFilePath, err))
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	z := zip.NewWriter(f)
	defer func() {
		if err := z.Close(); err != nil {
			panic(err)
		}
	}()

	skipMimetypeFile := false

	var addFileToZip = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the path of th e file relative to the folder we're zipping
		relativePath, err := filepath.Rel(tempDir, path)
		relativePath = filepath.ToSlash(relativePath)
		if err != nil {
			panic(fmt.Sprintf("Error closing EUB file: %s", err))
		}

		// Don't include directories
		if !info.Mode().IsRegular() {
			return nil
		}

		var w io.Writer
		if path == filepath.Join(tempDir, mimetypeFilename) {

			// Skip the mimetype file if it's already been written
			if skipMimetypeFile == true {
				return nil
			}

			// The mimetype file must be uncompressed according to EPUB spec
			w, err = z.CreateHeader(&zip.FileHeader{
				Name:   relativePath,
				Method: zip.Store,
			})
		} else {
			w, err = z.Create(relativePath)
		}
		if err != nil {
			panic(fmt.Sprintf("Error creating zip writer: %s", err))
		}

		r, err := os.Open(path)
		if err != nil {
			panic(fmt.Sprintf("Error opening file being added to EPUB: %s", err))
		}
		defer func() {
			if err := r.Close(); err != nil {
				panic(err)
			}
		}()

		_, err = io.Copy(w, r)
		if err != nil {
			panic(fmt.Sprintf("Error copying contents of the file being added to EPUB: %s", err))
		}

		return nil
	}

	// Add mimetype file first
	mimetypeFilePath := filepath.Join(tempDir, mimetypeFilename)
	mimetypeInfo, err := os.Lstat(mimetypeFilePath)
	if err != nil {
		panic(fmt.Sprintf("Unable to get FileInfo for mimetype file: %s", err))
	}
	err = addFileToZip(mimetypeFilePath, mimetypeInfo, nil)
	if err != nil {
		panic(fmt.Sprintf("Unable to add mimetype file to EPUB: %s", err))
	}

	skipMimetypeFile = true

	err = filepath.Walk(tempDir, addFileToZip)
	if err != nil {
		panic(fmt.Sprintf("Unable to add file to EPUB: %s", err))
	}

	return nil

}
