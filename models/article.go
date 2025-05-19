package models

import "time"

type Article struct {
	Connection  string    `connection:"sqlite"`
	ID          int64     `column:"id" type:"primary_key"`
	NewsQueryID int64     `column:"news_query_id" type:"INTEGER"`
	CreatedAt   time.Time `column:"created_at" type:"datetime_current"`
	Publisher   string    `json:"publisher" column:"publisher" type:"string"`
	PublishedAt time.Time `json:"published_at" column:"published_at" type:"datetime"`
	Image       string    `json:"image" column:"image" type:"string"`
	Title       string    `json:"title" column:"title" type:"string"`
	Description string    `json:"description" column:"description" type:"string"`
	Content     string    `json:"content" column:"content" type:"text"`
	Url         string    `json:"url" column:"url" type:"string"`
}
