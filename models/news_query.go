package models

import "time"

type NewsQuery struct {
	Connection string    `connection:"sqlite" table:"news_queries"`
	ID         int64     `json:"id" column:"id" type:"primary_key"`
	CreatedAt  time.Time `json:"created_at" column:"created_at" type:"datetime_current"`
	Url        string    `json:"url" column:"url" type:"string"`
}

type NewsQueryArticles struct {
	Connection  string    `connection:"sqlite" table:"news_queries"`
	ID          int64     `json:"id" column:"id" type:"primary_key"`
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
