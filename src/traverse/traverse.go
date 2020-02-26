package traverse

import (
	"os"
	"path/filepath"
	"strings"
)

// WalkFunc is the type of the function called for each file visited by WalkFiles.
type WalkFunc func(path, trimmed string, info os.FileInfo) error

// WalkFiles walks the file tree rooted at root, calling walkFn for each file.
func WalkFiles(folder string, walkFn WalkFunc) error {
	folderPrefix := filepath.Clean(folder)

	if folderPrefix == "." {
		folderPrefix = ""
	}

	return filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		trimmed := strings.TrimPrefix(path, folderPrefix)
		trimmed = strings.TrimPrefix(trimmed, "/")

		return walkFn(path, trimmed, info)
	})
}
