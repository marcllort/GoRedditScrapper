package main

import (
	"DataRetriever/downloadHelper"
	"DataRetriever/utils"
	"fmt"
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
			downloadHelper.DownloadFile(tempPost.StoryURL, tempPost.Title, utils.ImagePath(redditRepo))
		} else if strings.Contains(tempPost.StoryURL, "imgur") {
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
