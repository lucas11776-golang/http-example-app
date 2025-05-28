package models

import "time"

type Client struct {
	Connection string    `connection:"sqlite"`
	ID         int64     `column:"id" type:"primary_key"`
	CreatedAt  time.Time `column:"created_at" type:"datetime_current"`
}
