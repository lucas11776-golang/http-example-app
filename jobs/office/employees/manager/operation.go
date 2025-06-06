package manager

import (
	"context"
	"fmt"
	"server/jobs/workspace"
	"server/jobs/workspace/paperwork/analyst"
	"server/models"
)

// TODO: Operation manager will deal with client interests
type OperationManager struct {
	workspace *workspace.Workspace
}

// Comment
func NewOperationManager(workspace *workspace.Workspace) *OperationManager {
	return &OperationManager{
		workspace: workspace,
	}
}

// Comment
func (ctx *OperationManager) Work(context context.Context) {
	// Do some work
}

// Comment
func (ctx *OperationManager) PublishArticles(context context.Context, verifiedArticles []*analyst.ArticleVerified) ([]*models.Article, error) {
	fmt.Println("Publishing Articles", verifiedArticles)

	return nil, nil
}

// Comment
func (ctx *OperationManager) PublishArticle(context context.Context, verifiedArticle *analyst.ArticleVerified) (*models.Article, error) {
	fmt.Println("Publishing Article", verifiedArticle)

	return nil, nil
}
