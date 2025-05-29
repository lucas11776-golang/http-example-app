package designer

import (
	"fmt"
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
	ctx.activities()

	for range time.Tick(time.Minute * 10) {
		ctx.activities()
	}
}

// Comment
func (ctx *GraphicDesigner) activities() {
	fmt.Printf("%s is working hard\r\n", "GraphicDesigner")
}
