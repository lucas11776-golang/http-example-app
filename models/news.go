package models

import "time"

type Article struct {
	Publisher   string    `json:"publisher"`
	PublishedAt time.Time `json:"published_at"`
	Image       string    `json:"image"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Url         string    `json:"url"`
}
