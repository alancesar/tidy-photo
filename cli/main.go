package main

import (
	"flag"
	"fmt"
	"github.com/alancesar/tidy-file/command"
	"github.com/alancesar/tidy-file/mime"
	"github.com/alancesar/tidy-file/path"
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
	// Some RAW image files as Canon's .CR3 are described as "application/octet-stream"
	paths := path.LookFor(*rootSourcePath, mime.ImageType, mime.ApplicationOctetStream)
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
	raw, err := exif.NewReader(sourcePath).Extract()
	if err != nil {
		return "", err
	}

	parser := exif.NewParser(raw)
	t, err := parser.GetDateTime()
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
