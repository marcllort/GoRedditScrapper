package downloadHelper

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(URL, fileName string, directory string) error {
	//Get the response bytes from the url

	response, err := http.Get(URL)
	if err != nil {
	}
	defer response.Body.Close()

	fileName = fileName + ".png"
	path, err := os.Getwd()
	filePath := filepath.Join(path, directory)

	os.MkdirAll(filePath, 0644)

	//Create a empty file
	file, err := os.Create(filepath.Join(filePath, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
