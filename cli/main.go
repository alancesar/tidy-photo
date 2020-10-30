package main

import (
	"flag"
	"fmt"
	"github.com/alancesar/tidy-photo/command"
	"github.com/alancesar/tidy-photo/file"
	"github.com/alancesar/tidy-photo/processor"
	"os"
	"path/filepath"
)

const defaultPattern = "2006/2006-01-02"

func main() {
	rootSourcePath := flag.String("s", "./", "source directory")
	rootDestinationPath := flag.String("o", "./", "output directory")
	pattern := flag.String("p", defaultPattern, "output pattern")
	sandbox := flag.Bool("t", false, "run in test mode")
	flag.Parse()

	fmt.Println("Reading source directory...")
	paths := make([]string, 0)
	_ = filepath.Walk(*rootSourcePath, func(path string, info os.FileInfo, err error) error {
		if !file.IsFile(path) {
			return nil
		}

		paths = append(paths, path)
		return nil
	})

	total := len(paths)

	var commands []command.Command
	if !*sandbox {
		commands = []command.Command{command.MkDirCommand, os.Rename}
	}

	for index, path := range paths {
		destination, err := processor.Process(path, *rootDestinationPath, *pattern, commands...)
		if err != nil && err != processor.ExifErr {
			panic(err)
		}

		if err == nil {
			fmt.Printf("(%d/%d) %s\n", index+1, total, destination)
		}
	}
}
