package path

import (
	"github.com/alancesar/tidy-photo/mime"
	"os"
	"path/filepath"
)

// LookFor deep walks in a path and get all files that match with a provided mime.Type.
func LookFor(rootPath string, t mime.Type) []string {
	paths := make([]string, 0)
	_ = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() && info.Mode().IsRegular() && mime.Is(path, t) {
			paths = append(paths, path)
		}

		return nil
	})

	return paths
}
