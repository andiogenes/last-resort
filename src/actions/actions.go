package actions

import (
	"github.com/andiogenes/last-resort/src/translator"
	"github.com/andiogenes/last-resort/src/traverse"
	"io/ioutil"
	"os"
	"path/filepath"
)

// TranslateFile reads data from source, translates it with given lexeme and writes new data in destination file.
func TranslateFile(source, destination string, lexeme []byte) error {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	translated := translator.TranslateBytes(input, lexeme)

	return ioutil.WriteFile(destination, translated, 0644)
}

// TranslateFolder walks over given source folder, translates files in it and saves them in given destination folder.
func TranslateFolder(source, destination string, lexeme []byte) error {
	return traverse.WalkFiles(source, func(path, trimmed string, info os.FileInfo) error {
		absoluteDestination := filepath.Join(destination, trimmed)

		// Create folder if current file in traverse is directory
		if info.IsDir() {
			_ = os.Mkdir(absoluteDestination, os.ModePerm)
			return nil
		}

		return TranslateFile(path, absoluteDestination, lexeme)
	})
}