package utils

import (
	"os"
	"path/filepath"
)

// GetCurrentDirName returns name of current directory.
func GetCurrentDirName() (string, error) {
	path, err := filepath.Abs(".")
	if err != nil {
		return "", err
	}

	return filepath.Base(path), nil
}

// IsDir returns true if file with the given name is directory, false otherwise.
func IsDir(fileName string) (bool, error) {
	info, err := os.Stat(fileName)
	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}