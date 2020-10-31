package exif

import (
	exif2 "github.com/dsoprea/go-exif/v3"
	"os"
)

// Reader is the structure responsible for extract raw EXIF data.
type Reader struct {
	path string
}

// NewReader creates a new *Reader.
func NewReader(path string) *Reader {
	return &Reader{
		path: path,
	}
}

// Extract extracts the raw exif from provided path.
func (e *Reader) Extract() ([]byte, error) {
	source, err := os.Open(e.path)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = source.Close()
	}()

	return exif2.SearchAndExtractExifWithReader(source)
}
