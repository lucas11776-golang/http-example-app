package junior

import (
	"fmt"
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
	ctx.activities()

	for range time.Tick(time.Minute * 10) {
		ctx.activities()
	}
}

// Comment
func (ctx *JuniorAnalyst) activities() {
	go ctx.researchNews()
}

// Comment
func (ctx *JuniorAnalyst) researchNews() {
	fmt.Printf("%s is working hard\r\n", "JuniorAnalyst")
}
