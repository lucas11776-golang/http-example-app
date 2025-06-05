package manager

import (
	"context"
	"server/jobs/workspace"
)

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
}
