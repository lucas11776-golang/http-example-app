package office

import (
	"context"
	"server/models"
)

type OperationManager interface {
	PublishArticle(context context.Context, article models.ArticleCaputure)
	PublishClientArticle(context context.Context, client models.Client, article models.ArticleCaputure)
}

type SeniorAnalyst interface {
	ArticleValidation()
	ArticleImageDescription()
}

type JuniorAnalyst interface {
	Articles(context context.Context)
	ArticlesByClientIntrests(context context.Context, client *models.Client)
}

type GraphicDesigner interface {
	DesignImageConcept()
}
