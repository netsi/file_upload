package image

import (
	"net/http"
	"strings"
)

var supportedTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}

// GetSupportedTypes returns the list of the supportedTypes in a comma separated string.
func GetSupportedTypes() string {
	s := []string{}
	for supportedType := range supportedTypes {
		s = append(s, supportedType)
	}

	return strings.Join(s, ",")
}

// IsSupportedImage returns true if the file type detected is one of the supportedTypes.
func IsSupportedImage(file []byte) bool {
	fileType := http.DetectContentType(file)
	_, ok := supportedTypes[fileType]

	return ok
}
