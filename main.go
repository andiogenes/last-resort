package main

import (
	"flag"
	"fmt"
	"github.com/andiogenes/last-resort/src/actions"
	"github.com/andiogenes/last-resort/src/utils"
	"log"
	"os"
	"path/filepath"
)

// parseFlags parses command-line flags and returns values in "string->string" map.
func parseFlags() map[string]string {
	outputFlag := flag.String("o", "", "")
	lexemeFlag := flag.String("l", "", "")

	flag.Parse()

	return map[string]string{
		"o": *outputFlag,
		"l": *lexemeFlag,
	}
}

// resolveOutputFileName returns default output file name if real is empty.
func resolveOutputFileName(fileName string, pattern string) string {
	if fileName == "" {
		return pattern+".fk"
	}

	return fileName
}

// resolveOutputFolderName returns default output folder name if real is empty.
func resolveOutputFolderName(folderName string, pattern string) string {
	if folderName == "" {
		return filepath.Join(pattern, "..", pattern+"-fk")
	}

	return folderName
}

func main() {
	// Get information from application environment
	flags := parseFlags()
	arg := flag.Arg(0)

	if arg == "" {
		fmt.Println("[arg] shouldn't be empty")
		os.Exit(2)
	}

	isDirectory, err := utils.IsDir(arg)
	if err != nil {
		log.Fatal(err)
	}

	// Perform actions relatively to passed flags and arguments.
	if isDirectory {
		err = actions.TranslateFolder(arg, resolveOutputFolderName(flags["o"], arg), []byte(flags["l"]))
	} else {
		err = actions.TranslateFile(arg, resolveOutputFileName(flags["o"], arg), []byte(flags["l"]))
	}

	if err != nil {
		log.Fatal(err)
	}
}
