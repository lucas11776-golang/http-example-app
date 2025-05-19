package models

import "time"

type NewsQuery struct {
	Connection string    `connection:"sqlite" table:"news_queries"`
	ID         int64     `json:"id" column:"id" type:"primary_key"`
	CreatedAt  time.Time `json:"created_at" column:"created_at" type:"datetime_current"`
	Url        string    `json:"first_name" column:"url" type:"string"`
}
