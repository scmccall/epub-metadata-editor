package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//  code from from https://golangcode.com/unzip-files-in-go/

func Unzip(src string, ext string, dest string) ([]string, error) {

	file := src + ext

	var filenames []string

	r, err := zip.OpenReader(file)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}

	// Remove zip file so it can be recreated later
	os.Remove(file)

	return filenames, nil
}

func Zip(filename string, ext string) error {
	// Creates .epub file
	file, err := os.Create(filename + ext)
	if err != nil {
		log.Fatal("os.Create(filename) error: ", err)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		fmt.Println("Crawling: " + path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		switch {
		case info.IsDir():
			// f, err := w.Create(path + "/")
			f, err := w.Create(path)
			if err != nil {
				return err
			}
			_, err = io.Copy(f, file)
			if err != nil {
				return err
			}
		default:
			f, err := w.Create(path)
			if err != nil {
				return err
			}
			_, err = io.Copy(f, file)
			if err != nil {
				return err
			}
		}

		// f, err := w.Create(path)
		// if err != nil {
		// 	return err
		// }

		// _, err = io.Copy(f, file)
		// if err != nil {
		// 	return err
		// }

		return nil
	}

	func Zipdraft(filename string, ext string) error {
		file, err := os.Create(filename + ext)
		if err != nil {
			log.Fatal("os.Create(filename) error: ", err)
		}
		defer file.Close()

		w := zip.NewWriter(file)
		defer w.Close()

		walker := func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			relativePath, err := filepath.Rel(filename, path)
			relativePath = filepath.ToSlash(relativePath)
			if err != nil {
				panic(fmt.Sprintf("Error clsing .zip file: %s", err))
			}

			// Only include regular files and not directories
			if !info.Mode().IsRegular() {
				return nil
			}
			
		}
	}

	err = filepath.Walk(filename, walker)
	if err != nil {
		log.Fatal("filepath.Walk error: ", err)
	}
	return err
}
