package exif

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exif2 "github.com/dsoprea/go-exif/v3"
	common "github.com/dsoprea/go-exif/v3/common"
	"time"
)

const (
	dateTimeTagName = "DateTime"
	dateTimeLayout  = "2006:01:02 15:04:05"
)

// Parser is the structure responsible for parse a raw EXIF data.
type Parser struct {
	rawExif []byte
}

// NewParser creates a new *Parser
func NewParser(rawExif []byte) *Parser {
	return &Parser{
		rawExif: rawExif,
	}
}

// GetExifTags return all exif tags contained in raw EXIF.
func (p *Parser) GetExifTags() ([]exif2.ExifTag, error) {
	tags, _, err := exif2.GetFlatExifData(p.rawExif, nil)
	return tags, err
}

// GetThumbnail return the contained thumbnail in raw EXIF.
func (p *Parser) GetThumbnail() ([]byte, error) {
	im, err := common.NewIfdMappingWithStandard()
	if err != nil {
		return nil, err
	}

	ti := exif.NewTagIndex()

	_, index, err := exif.Collect(im, ti, p.rawExif)
	if err != nil {
		return nil, err
	}

	var thumbnail []byte
	if ifd, found := index.Lookup[exif2.ThumbnailFqIfdPath]; found == true {
		thumbnail, err = ifd.Thumbnail()
		if err != nil && err != exif2.ErrNoThumbnail {
			return nil, err
		}
	}

	return thumbnail, nil
}

// GetDateTime retrieves a time.Time.
func (p *Parser) GetDateTime() (time.Time, error) {
	tags, _, err := exif2.GetFlatExifData(p.rawExif, nil)
	if err != nil {
		return time.Now(), err
	}

	for _, item := range tags {

		if item.TagName == dateTimeTagName {
			t, err := time.Parse(dateTimeLayout, fmt.Sprint(item.Value))
			if err != nil {
				return time.Now(), err
			}

			return t, nil
		}
	}

	return time.Now(), nil
}
