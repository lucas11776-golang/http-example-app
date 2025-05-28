package models

import "time"

// TODO: If we have many instance running working we do not need to repeat active
type NewsJob struct {
	Connection string    `connection:"sqlite"`
	ID         int64     `column:"id" type:"primary_key"`
	CreatedAt  time.Time `column:"created_at" type:"datetime_current"`
}

// TODO: If we have many instance running working we do not need to repeat active
type NewsClientJob struct {
	Connection string    `connection:"sqlite"`
	ID         int64     `column:"id" type:"primary_key"`
	CreatedAt  time.Time `column:"created_at" type:"datetime_current"`
}
