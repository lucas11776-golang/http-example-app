package office

import (
	"context"
	"server/models"
)

type SeniorAnalyst interface {
	ArticleValidate()
}

type OperationManager interface {
	PublishArticle(context context.Context, article models.ArticleCaputure)
	PublishClientArticle(context context.Context, client models.Client, article models.ArticleCaputure)
}

type JuniorAnalyst interface {
	Articles(context context.Context)
	ArticlesByClientIntrests(context context.Context, client *models.Client)
}

type Office struct {
	SeniorAnalyst SeniorAnalyst
	JuniorAnalyst JuniorAnalyst
}

// TODO: Whats happing in the office 2
func (ctx *Office) OfficeActive1() {}

// TODO: Whats happing in the office 2
func (ctx *Office) OfficeActive2() {}
