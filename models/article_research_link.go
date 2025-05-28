package models

import (
	"time"
)

type ArticleCaputure struct {
	Connection  string    `connection:"sqlite"`
	ID          int64     `column:"id" type:"primary_key"`
	CreatedAt   time.Time `column:"created_at" type:"datetime_current"`
	Source      string    `json:"website" column:"website" type:"string"`
	Title       string    `json:"title" column:"title" type:"string"`
	Image       string    `json:"image" column:"image" type:"string"`
	Category    string    `json:"category" column:"category" type:"string"`
	Description string    `json:"description" column:"description" type:"string"`
	Publisher   string    `json:"publisher" column:"publisher" type:"string"`
	PublishedAt time.Time `json:"published_at" column:"published_at" type:"datetime"`
	Content     time.Time `json:"content" column:"content" type:"text"`
}
