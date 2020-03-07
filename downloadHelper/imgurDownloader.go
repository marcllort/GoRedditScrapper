package downloadHelper

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func createFile(imageURL string, fileName string, directory string) {
	resp, err := http.Get(imageURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fileName = fileName + ".png"
	path, _ := os.Getwd()
	filePath := filepath.Join(path, directory)

	os.MkdirAll(filePath, 0644)

	//Create a empty file
	file, err := os.Create(filepath.Join(filePath, fileName))

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

}

func getVal(tag html.Token, fileName string, directory string) {
	// Finds that id value of the imgur links
	for _, div := range tag.Attr {
		if div.Key == "id" {
			if len(div.Val) == 7 {
				image := div.Val
				encodeURL(image, fileName, directory)
			}
		}
	}
}

func encodeURL(image string, fileName string, directory string) {
	imageURL := "https://i.imgur.com/" + image + ".jpg"
	createFile(imageURL, fileName, directory)
}

func ImgurDownload(url string, fileName string, directory string) {
	resp, _ := http.Get(url)
	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// Error Token is end of the document
			return
		case tt == html.StartTagToken:
			tag := z.Token()

			// Check to see if the tag has <div> if not move to the next line.
			div := tag.Data == "div"
			if !div {
				continue
			}

			getVal(tag, fileName, directory)
		}
	}
}
