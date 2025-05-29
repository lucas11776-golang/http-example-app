package senior

import (
	"fmt"
	"server/jobs/office/workspace"
	"time"
)

type SeniorAnalyst struct {
	workspace *workspace.Workspace
}

// Comment
func NewSeniorAnalyst(workspace *workspace.Workspace) *SeniorAnalyst {
	return &SeniorAnalyst{
		workspace: workspace,
	}
}

// Comment
func (ctx *SeniorAnalyst) Work() {
	ctx.activities()

	for range time.Tick(time.Minute * 10) {
		ctx.activities()
	}
}

// Comment
func (ctx *SeniorAnalyst) activities() {
	fmt.Printf("%s is working hard\r\n", "SeniorAnalyst")
}
