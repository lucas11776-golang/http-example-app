package employees

import (
	"context"
	"server/jobs/workspace/paperwork/analyst"
	"server/jobs/workspace/paperwork/designer"
	"server/models"
)

type OperationManager interface {
	Work(context context.Context)
	PublishArticles(context context.Context, verifiedArticles []*analyst.ArticleVerified) ([]*models.Article, error)
	PublishArticle(context context.Context, verifiedArticles *analyst.ArticleVerified) (*models.Article, error)
}

type SeniorAnalyst interface {
	Work(context context.Context)
	ResearchArticles(context context.Context, interest []string) ([]*analyst.ArticleVerified, error)
	VerifiedArticle(context context.Context, unverifiedArticles *analyst.ArticleCapture) (*analyst.ArticleVerified, error)
}

type JuniorAnalyst interface {
	Work(context context.Context)
	ResearchArticles(context context.Context, interest []string) ([]*analyst.ArticleCapture, error)
	UnverifiedArticles(context context.Context) ([]*analyst.ArticleCapture, error)
}

type SeniorGraphicDesigner interface {
	Work(context context.Context)
	DesignArticleImage(context context.Context, verifiedArticles *analyst.ArticleVerified) (*designer.Image, error)
}
