package manager

import (
	"server/jobs/workspace"
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

}
