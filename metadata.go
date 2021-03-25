package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Takes content.opf file from an unzipped EPUb and removes all lines
// containing the <dc:subject> tag
func editMetadata(absPathToEpubDir string) {

	contentFile := absPathToEpubDir + "/content.opf"

	// Read file
	input, err := ioutil.ReadFile(contentFile)
	if err != nil {
		panic(fmt.Sprintf("Error reading content.opf file: %s", err))
	}

	// Split file into array of stings
	lines := strings.Split(string(input), "\n")

	// Parse array for tag metadata, replace with empty string
	for i, line := range lines {
		if strings.Contains(line, "<dc:subject>") {
			lines[i] = ""
		}
	}

	// Remove empty lines
	output := strings.Join(lines, "\n")
	output = strings.Replace(output, "\n\n", "\n", -1)

	// Write to new content.opf file
	err = ioutil.WriteFile(contentFile, []byte(output), 0644)
	if err != nil {
		panic(fmt.Sprintf("Error writing content.opf file: %s", err))
	}

}
