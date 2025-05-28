package models

import "time"

type Article struct {
	Connection  string    `connection:"sqlite"`
	ID          int64     `column:"id" type:"primary_key"`
	NewsQueryID int64     `column:"news_query_id" type:"INTEGER"` // TODO: remove we are not using third party news...
	CreatedAt   time.Time `column:"created_at" type:"datetime_current"`
	Publisher   string    `json:"publisher" column:"publisher" type:"string"`
	PublishedAt time.Time `json:"published_at" column:"published_at" type:"datetime"`
	Image       string    `json:"image" column:"image" type:"string"`
	Title       string    `json:"title" column:"title" type:"string"`
	Description string    `json:"description" column:"description" type:"string"`
	Content     string    `json:"content" column:"content" type:"text"`
	Url         string    `json:"url" column:"url" type:"string"`
	Verified    bool      `json:"verified" column:"url" type:"boolean"`
}

type ArticleReport struct {
	Connection      string `connection:"sqlite" table:"article_reports"`
	ID              int64  `column:"id" type:"primary_key"`
	ArticleID       int64  `column:"article_id" type:"INTEGER"`
	NumberOfArticle int64  `column:"number_of_article" type:"INTEGER"`
}

type ArticeSource struct {
	Connection      string `connection:"sqlite" table:"article_reports"`
	ID              int64  `column:"id" type:"primary_key"`
	ArticleReportID int64  `column:"article_report_id" type:"INTEGER"`
}
