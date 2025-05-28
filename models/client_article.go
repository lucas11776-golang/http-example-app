package models

import "time"

// TODO: Type of article the client is looking.
type ClientArticleIntrest struct {
	Connection string    `connection:"sqlite"`
	ID         int64     `column:"id" type:"primary_key"`
	CreatedAt  time.Time `column:"created_at" type:"datetime_current"`
}

// TODO: Client Description of article in feed.
type ClientArticleDescritpion struct {
	Connection             string    `connection:"sqlite"`
	ID                     int64     `column:"id" type:"primary_key"`
	CreatedAt              time.Time `column:"created_at" type:"datetime_current"`
	ClientArticleIntrestID int64     `column:"client_article_intrest_id" type:"integer"`
	CategoryID             int64     `column:"category_id" type:"integer"`
}

// TODO: Type client articles relationship.
type ClientArticle struct {
	Connection string    `connection:"sqlite"`
	ID         int64     `column:"id" type:"primary_key"`
	CreatedAt  time.Time `column:"created_at" type:"datetime_current"`
	ClientID   int64     `column:"client_id" type:"integer"`
	ArticleID  int64     `column:"article_id" type:"integer"`
}
