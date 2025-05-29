package manager

import (
	"fmt"
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
	ctx.activities()

	for range time.Tick(time.Minute * 10) {
		ctx.activities()
	}
}

// Comment
func (ctx *OperationManager) activities() {
	fmt.Printf("%s is working hard\r\n", "OperationManager")
}
