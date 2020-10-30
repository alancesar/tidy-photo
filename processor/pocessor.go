package processor

import (
	"errors"
	"github.com/alancesar/tidy-photo/datetime"
	exif2 "github.com/dsoprea/go-exif/v3"
	"os"
	"path/filepath"
	"strings"
)

var ExifErr = errors.New("error on get exif")

func Process(sourcePath, rootDestination, pattern string, sandbox bool) (string, error) {
	source, err := os.Open(sourcePath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = source.Close()
	}()

	rawExif, err := exif2.SearchAndExtractExifWithReader(source)
	if err != nil {
		return "", err
	}

	tags, _, err := exif2.GetFlatExifData(rawExif, nil)
	if err != nil {
		return "", err
	}

	t, err := datetime.ExtractDateTime(tags)
	if err != nil {
		return "", err
	}

	date := t.Format(pattern)
	_, filename := filepath.Split(sourcePath)
	completePath := filepath.Join(strings.Split(date, "/")...)
	completePath = filepath.Join(rootDestination, completePath, filename)
	completePath = filepath.Clean(completePath)

	if !sandbox {
		dir, _ := filepath.Split(completePath)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", err
		}

		err = os.Rename(sourcePath, completePath)
	}

	return completePath, err
}
