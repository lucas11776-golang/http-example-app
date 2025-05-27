package models

import (
	"server/utils/slices"
	"time"

	"github.com/lucas11776-golang/orm"
)

type ArticleResearch struct {
	Connection  string    `connection:"sqlite"`
	ID          int64     `column:"id" type:"primary_key"`
	CreatedAt   time.Time `column:"created_at" type:"datetime_current"`
	Url         string    `json:"url" column:"url" type:"string"`
	Title       string    `json:"title" column:"title" type:"string"`
	Category    string    `json:"category" column:"category" type:"string"`
	Description string    `json:"description" column:"description" type:"string"`
	Publisher   string    `json:"publisher" column:"publisher" type:"string"`
	PublishedAt time.Time `json:"published_at" column:"published_at" type:"datetime"`
}

// Comment
func (ctx *ArticleResearch) Save() (*ArticleResearch, error) {
	data, err := slices.StructToMap(ctx)

	if err != nil {
		return nil, err
	}

	return orm.Model(ArticleResearch{}).Insert(data)
}
