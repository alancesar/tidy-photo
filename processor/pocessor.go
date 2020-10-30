package processor

import (
	"github.com/alancesar/tidy-photo/command"
	"github.com/alancesar/tidy-photo/datetime"
	"github.com/alancesar/tidy-photo/exif"
	"path/filepath"
	"strings"
)

func Process(sourcePath, rootDestination, pattern string, commands ...command.Command) (string, error) {
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
	destinationPath := filepath.Join(strings.Split(date, "/")...)
	destinationPath = filepath.Join(rootDestination, destinationPath, filename)
	destinationPath = filepath.Clean(destinationPath)

	err = command.NewExecutor(sourcePath, destinationPath).Execute(commands...)
	return destinationPath, err
}
