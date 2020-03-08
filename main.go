package main

import (
	"DataRetriever/database"
	"DataRetriever/utils"
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func retrieveReddit(reddits []string) {
	db := database.CreateConnection()
	posts := utils.ScrapPosts(reddits, db)
	fmt.Println(posts)
}

func main() {

	reddits := []string{"https://old.reddit.com/r/funny", "https://old.reddit.com/r/tinder"}

	gocron.Every(15).Seconds().DoSafely(retrieveReddit, reddits)
	<-gocron.Start()

	// TODO: PARAMETERS, to specify where to download, crontab times, config file location?, number of downloads, links to download from...

}
