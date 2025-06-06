package analyst

import (
	"time"
)

type ArticleVerifiedSource struct {
	Connection        string    `connection:"sqlite" table:"article_verified_sources"`
	ID                int64     `column:"id" type:"primary_key"`
	CreatedAt         time.Time `column:"created_at" type:"datetime_current"`
	ArticleVerifiedID string    `column:"article_verified_id" type:"string"`
	Name              string    `column:"name" type:"string"`
	Website           string    `column:"website" type:"string"`
	Trusted           string    `column:"trusted" type:"boolean"`
}

type ArticleVerified struct {
	Connection       string    `connection:"sqlite" table:"verified_articles"`
	ID               int64     `column:"id" type:"primary_key"`
	ArticleCaptureID int64     `column:"article_capture_id" type:"integer"`
	CreatedAt        time.Time `column:"created_at" type:"datetime_current"`
	PublishedAt      time.Time `column:"published_at" type:"datetime"`
	Rating           string    `column:"rating" type:"integer"`
	Title            string    `column:"title" type:"string"`
	Image            string    `column:"image" type:"string"`
	Description      string    `column:"description" type:"string"`
	Content          string    `column:"content" type:"text"`
	Html             string    `column:"html" type:"text"`
	Trusted          []*ArticleVerifiedSource
	Untrusted        []*ArticleVerifiedSource
}
