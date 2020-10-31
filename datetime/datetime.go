package datetime

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	"time"
)

const (
	dateTimeTagName = "DateTime"
	dateTimeLayout  = "2006:01:02 15:04:05"
)

// ExtractDateTime retrieves a time.Time from provided exif tags.
func ExtractDateTime(tags []exif.ExifTag) (time.Time, error) {
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
