package analyst

import (
	"time"

	"github.com/lucas11776-golang/orm"
)

type ArticleCapture struct {
	Connection     string    `connection:"sqlite" table:"article_captures"`
	ID             int64     `column:"id" type:"primary_key"`
	CreatedAt      time.Time `column:"created_at" type:"datetime_current"`
	VerificationAt time.Time `column:"verification_at" type:"datetime"`
	VerifiedAt     time.Time `column:"verified_at" type:"datetime"`
	PublishedAt    time.Time `column:"published_at" type:"date"`
	Publisher      string    `column:"publisher" type:"string"`
	Image          string    `column:"image" type:"string"`
	Title          string    `column:"title" type:"string"`
	Category       string    `column:"category" type:"string"`
	Website        string    `column:"website" type:"string"`
	Description    string    `column:"description" type:"string"`
	Content        string    `column:"content" type:"text"`
}

// Comment
func (ctx *ArticleCapture) Verified(t time.Time) error {
	return orm.Model(ArticleCapture{}).
		Update(orm.Values{"verified_at": t.Format(time.DateTime)})
}

// Comment
func (ctx *ArticleCapture) Verifying(t time.Time) error {
	return orm.Model(ArticleCapture{}).
		Update(orm.Values{"verification_at": t.Format(time.DateTime)})
}
