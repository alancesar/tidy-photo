package exif

import (
	exif2 "github.com/dsoprea/go-exif/v3"
	"os"
)

// Extractor is the structure responsible for extract exif tags.
type Extractor struct {
	path string
}

// NewExtractor creates a new *Extractor.
func NewExtractor(path string) *Extractor {
	return &Extractor{
		path: path,
	}
}

// Extract extracts all exif tags from provided path.
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
