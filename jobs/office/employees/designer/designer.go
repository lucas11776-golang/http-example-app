package designer

import (
	"context"
	"server/jobs/workspace"
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
