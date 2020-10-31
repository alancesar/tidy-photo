package main

import (
	"flag"
	"fmt"
	"github.com/alancesar/tidy-file/command"
	"github.com/alancesar/tidy-file/mime"
	"github.com/alancesar/tidy-file/path"
	"github.com/alancesar/tidy-photo/datetime"
	"github.com/alancesar/tidy-photo/exif"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultPattern   = "2006/2006-01-02"
	defaultSeparator = "/"
)

func main() {
	rootSourcePath := flag.String("s", "./", "source directory")
	rootDestinationPath := flag.String("o", "./", "output directory")
	pattern := flag.String("p", defaultPattern, "output pattern")
	sandbox := flag.Bool("t", false, "run in test mode")
	flag.Parse()

	fmt.Println("Reading source directory...")
	paths := path.LookFor(*rootSourcePath, mime.ImageType)
	total := len(paths)

	var commands []command.Command
	if !*sandbox {
		commands = []command.Command{command.MkDirCommand, os.Rename}
	}

	for index, sourcePath := range paths {
		destination, err := process(sourcePath, *rootDestinationPath, *pattern, commands...)
		if err != nil {
			fmt.Printf("(%d/%d) [failed ] %s\n", index+1, total, destination)
		} else {
			fmt.Printf("(%d/%d) [success] %s\n", index+1, total, destination)
		}
	}
}

func process(sourcePath, rootDestination, pattern string, commands ...command.Command) (string, error) {
	tags, err := exif.NewExtractor(sourcePath).Extract()
	if err != nil {
		return "", err
	}

	t, err := datetime.ExtractDateTime(tags)
	if err != nil {
		return "", err
	}

	date := t.Format(pattern)
	_, filename := filepath.Split(sourcePath)
	destinationPath := filepath.Join(strings.Split(date, defaultSeparator)...)
	destinationPath = filepath.Join(rootDestination, destinationPath, filename)
	destinationPath = filepath.Clean(destinationPath)

	err = command.NewExecutor(sourcePath, destinationPath).Execute(commands...)
	return destinationPath, err
}
