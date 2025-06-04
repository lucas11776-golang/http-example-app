package manager

import (
	"server/jobs/office/workspace"
	"time"
)

type StudioManager struct {
	workspace *workspace.Workspace
}

// Comment
func NewStudioManager(workspace *workspace.Workspace) *StudioManager {
	return &StudioManager{
		workspace: workspace,
	}
}

// Comment
func (ctx *StudioManager) Work() {
	for range time.Tick(time.Minute * 10) {
	}
}
