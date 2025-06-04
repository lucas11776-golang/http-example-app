package designer

import (
	"server/jobs/office/workspace"
	"time"
)

type GraphicDesigner struct {
	workspace *workspace.Workspace
}

// Comment
func NewGraphicDesigner(workspace *workspace.Workspace) *GraphicDesigner {
	return &GraphicDesigner{
		workspace: workspace,
	}
}

// Comment
func (ctx *GraphicDesigner) Work() {
	for range time.Tick(time.Minute * 10) {
	}
}
