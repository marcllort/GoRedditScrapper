package model

import "time"

type Post struct {
	Title       string
	Subreddit   string
	Hash        []byte
	StoryURL    string
	Comments    string
	RetrievedAt time.Time
}
