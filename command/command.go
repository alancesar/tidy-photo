package command

import (
	"os"
	"path/filepath"
)

type Command func(source, destination string) error

func MkDirCommand(_, destination string) error {
	dir, _ := filepath.Split(destination)
	return os.MkdirAll(dir, os.ModePerm)
}
