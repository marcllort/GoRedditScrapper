package downloadHelper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
)

func GfycatDownload(url string, fileName string, directory string) {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if property, _ := s.Attr("property"); property == "og:video" {
			downloadURL, _ := s.Attr("content")
			fmt.Printf("Description field: %s\n", downloadURL)
			DownloadFile(downloadURL, fileName, directory, "mp4")
		}
	})
}
