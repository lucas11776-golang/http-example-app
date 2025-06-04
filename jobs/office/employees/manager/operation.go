package manager

import (
	"server/jobs/office/workspace"
	"time"
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
func (ctx *OperationManager) Work() {
	for range time.Tick(time.Minute * 10) {
	}
}
