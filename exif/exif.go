package exif

import (
	exif2 "github.com/dsoprea/go-exif/v3"
	"os"
)

type Extractor struct {
	path string
}

func NewExtractor(path string) *Extractor {
	return &Extractor{
		path: path,
	}
}

func (e *Extractor) Extract() ([]exif2.ExifTag, error) {
	source, err := os.Open(e.path)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = source.Close()
	}()

	rawExif, err := exif2.SearchAndExtractExifWithReader(source)
	if err != nil {
		return nil, err
	}

	tags, _, err := exif2.GetFlatExifData(rawExif, nil)
	return tags, err
}
