package main

import (
	"DataRetriever/downloadHelper"
	"DataRetriever/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type item struct {
	StoryURL    string
	Source      string
	RetrievedAt time.Time
	Comments    string
	Title       string
}

func main() {
	posts := []item{}
	redditRepo := ""

	// Instantiate default collector
	collector := colly.NewCollector(
		// Visit only domains: old.reddit.com
		colly.AllowedDomains("old.reddit.com"),
		colly.Async(true),
	)

	// On every a element which has .top-matter attribute call callback
	// This class is unique to the div that holds all information about a story
	collector.OnHTML(".top-matter", func(e *colly.HTMLElement) {
		tempPost := item{}
		tempPost.StoryURL = e.ChildAttr("a[data-event-action=title]", "href")
		tempPost.Source = redditRepo
		tempPost.Title = e.ChildText("a[data-event-action=title]")
		tempPost.Comments = e.ChildAttr("a[data-event-action=comments]", "href")
		tempPost.RetrievedAt = time.Now()
		if strings.HasSuffix(tempPost.StoryURL, "jpg") || strings.HasSuffix(tempPost.StoryURL, "png") {
			downloadFile(tempPost.StoryURL, tempPost.Title, utils.ImagePath(redditRepo))
		}
		if strings.Contains(tempPost.StoryURL, "imgur") {
			downloadHelper.ImgurDownload(tempPost.StoryURL, tempPost.Title, utils.UrlImagePath(e.Request.URL))
		}
		posts = append(posts, tempPost)
	})

	// On every span tag with the class next-button
	/*collector.OnHTML("span.next-button", func(h *colly.HTMLElement) {
		t := h.ChildAttr("a", "href")
		collector.Visit(t)
	})*/

	// Set max Parallelism and introduce a Random Delay
	collector.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	// Before making a request print "Visiting ..."
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())

	})

	// Crawl all reddits the user passes in
	reddits := [...]string{"https://old.reddit.com/r/funny", "https://old.reddit.com/r/tinder"}
	for _, reddit := range reddits {
		redditRepo = reddit
		collector.Visit(reddit)
	}

	collector.Wait()
	fmt.Println(posts)

}

func downloadFile(URL, fileName string, directory string) error {
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
