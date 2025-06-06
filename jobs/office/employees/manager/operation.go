package manager

import (
	"context"
	"server/jobs/workspace"
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
}
