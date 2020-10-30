package command

import (
	"os"
	"path/filepath"
)

type Command func(source, destination string) error

// MkDirCommand create a new directory if it does not exist.
func MkDirCommand(_, destination string) error {
	dir, _ := filepath.Split(destination)
	return os.MkdirAll(dir, os.ModePerm)
}
