package mime

import (
	"github.com/gabriel-vasile/mimetype"
	"strings"
)

type Type string

const (
	ImageType Type = "audio/"
)

func Is(path string, t Type) bool {
	mime, err := mimetype.DetectFile(path)
	if err != nil {
		return false
	}

	return strings.Contains(mime.String(), string(t))
}
