package junior

import (
	"server/jobs/office/workspace"
	"time"
)

type JuniorAnalyst struct {
	workspace *workspace.Workspace
}

// Comment
func NewJuniorAnalyst(workspace *workspace.Workspace) *JuniorAnalyst {
	return &JuniorAnalyst{
		workspace: workspace,
	}
}

// Comment
func (ctx *JuniorAnalyst) Work() {
	for range time.Tick(time.Minute * 10) {
	}
}
