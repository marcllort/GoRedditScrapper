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
	Title    string
	StoryURL string
	//dataURL     string
	//Source      string
	Comments    string
	RetrievedAt time.Time
}

func main() {
	posts := []item{}
	subreddit := ""

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
		tempPost.Title = e.ChildText("a[data-event-action=title]")
		tempPost.Comments = e.ChildAttr("a[data-event-action=comments]", "href")
		tempPost.RetrievedAt = time.Now()
		if strings.HasSuffix(tempPost.StoryURL, "jpg") || strings.HasSuffix(tempPost.StoryURL, "png") {
			downloadHelper.DownloadFile(tempPost.StoryURL, tempPost.Title, subreddit, "png")
		} else if strings.Contains(tempPost.StoryURL, "imgur") {
			downloadHelper.ImgurDownload(tempPost.StoryURL, tempPost.Title, subreddit)
		} else if strings.Contains(tempPost.StoryURL, "gfycat") {
			downloadHelper.GfycatDownload(tempPost.StoryURL, tempPost.Title, subreddit)
		} else {
			collector.Visit(tempPost.Comments)
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
		subreddit = utils.ImagePath(reddit)
		collector.Visit(reddit)
		time.Sleep(2 * time.Second)
	}

	collector.Wait()
	fmt.Println(posts)

}
