package util

import (
	"net/http"
	"regexp"
)

func IsValidFileExt(ext string) bool {
	validExt := []string{"jpg", "jpeg", "png"}
	valid := false
	for _, item := range validExt {
		if ext == item {
			valid = true
			break
		}
	}
	return valid
}

func IsImage(data []byte) bool {
	re, _ := regexp.Compile("^image")
	contentType := http.DetectContentType(data)
	valid := re.Match([]byte(contentType))
	return valid
}
