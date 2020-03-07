package utils

import (
	"net/url"
	"path"
	"path/filepath"
)

func ImagePath(url string) string {

	response := path.Base(url)
	filePath := filepath.Join("data", response)

	return filePath
}

func UrlImagePath(url *url.URL) string {

	response := path.Base(url.String())
	filePath := filepath.Join("data", response)

	return filePath
}
