package designer

import (
	"context"
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
}

// Comment
func (ctx *SeniorGraphicDesigner) DesignArticleImage(context context.Context, article analyst.ArticleVerified) (*designer.Image, error) {
	return nil, nil
}
