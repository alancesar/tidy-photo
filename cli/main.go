package main

import (
	"flag"
	"fmt"
	"github.com/alancesar/tidy-photo/command"
	"github.com/alancesar/tidy-photo/mime"
	"github.com/alancesar/tidy-photo/path"
	"github.com/alancesar/tidy-photo/processor"
	"os"
)

const defaultPattern = "2006/2006-01-02"

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
		destination, err := processor.Process(sourcePath, *rootDestinationPath, *pattern, commands...)
		if err != nil {
			fmt.Printf("(%d/%d) [failed ] %s\n", index+1, total, destination)
		} else {
			fmt.Printf("(%d/%d) [success] %s\n", index+1, total, destination)
		}
	}
}
