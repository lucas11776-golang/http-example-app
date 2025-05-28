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

type GraphicDesigner interface {
}

type Office struct {
	SeniorAnalyst   SeniorAnalyst
	JuniorAnalyst   JuniorAnalyst
	GraphicDesigner GraphicDesigner
}

// Comment
func (ctx *Office) LoadLatestNews() {

}

// Comment
func (ctx *Office) LoadLatestNewsClients() {

}
