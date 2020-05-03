package utils

import (
	"GoRedditScrapper/model"
	"log"
	"os/exec"
)

func ScrapPostsTikTok(reddits []string) []model.Post {
	posts := []model.Post{}

	// Only working on unix based systems
	//https://github.com/drawrowfly/tiktok-scraper#to-do
	cmd := exec.Command("bash", "tiktok-scraper", "trend -n 20 -d --filepath C:\\Users\\mac12\\Atollic")

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	return posts
}
