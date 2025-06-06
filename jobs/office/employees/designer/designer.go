package designer

import (
	"context"
	"fmt"
	"server/jobs/workspace"
	"server/jobs/workspace/paperwork/analyst"
	"server/jobs/workspace/paperwork/designer"
)

type SeniorGraphicDesigner struct {
	workspace *workspace.Workspace
}

// Comment
func NewSeniorGraphicDesigner(workspace *workspace.Workspace) *SeniorGraphicDesigner {
	return &SeniorGraphicDesigner{
		workspace: workspace,
	}
}

// Comment
func (ctx *SeniorGraphicDesigner) Work(context context.Context) {
	// Do some work
}

// Comment
func (ctx *SeniorGraphicDesigner) DesignArticleImage(context context.Context, verifiedArticles *analyst.ArticleVerified) (*designer.Image, error) {
	fmt.Println("Designing Article Image:", verifiedArticles)

	return nil, nil
}
