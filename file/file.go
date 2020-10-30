package file

import "os"

func IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	if stat.Mode().IsDir() {
		return false
	}

	if !stat.Mode().IsRegular() {
		return false
	}

	return true
}
