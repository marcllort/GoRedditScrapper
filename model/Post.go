package model

import "time"

type Post struct {
	Title    string
	StoryURL string
	//dataURL     string
	//Source      string
	Comments    string
	RetrievedAt time.Time
}
