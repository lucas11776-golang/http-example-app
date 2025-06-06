package office

import (
	"context"
	"fmt"
	"server/jobs/office/employees/analyst/junior"
	"server/jobs/office/employees/analyst/senior"
	"server/jobs/office/employees/designer"
	"server/jobs/office/employees/manager"
	"server/jobs/workspace"
	"time"
)

type Office struct {
	workspace *workspace.Workspace
}

// Comment
func NewOffice() *Office {
	workspace := workspace.NewWorkspace()

	workspace.OperationManager = manager.NewOperationManager(workspace)
	workspace.SeniorAnalyst = senior.NewSeniorAnalyst(workspace)
	workspace.SeniorGraphicDesigner = designer.NewSeniorGraphicDesigner(workspace)
	workspace.JuniorAnalyst = junior.NewJuniorAnalyst(workspace)

	return &Office{workspace: workspace}
}

// Comment
func (ctx *Office) Launch(context context.Context) {
	// ctx.duties(context)

	fmt.Println(ctx.workspace.SeniorAnalyst.ResearchArticles(context, []string{}))

	// fmt.Println(ctx.workspace.JuniorAnalyst.ResearchArticles(context, []string{}))

	for range time.Tick(time.Minute * 10) {
		ctx.duties(context)
	}
}

// Comment
func (ctx *Office) duties(context context.Context) {
	go ctx.workspace.JuniorAnalyst.Work(context)
	go ctx.workspace.SeniorAnalyst.Work(context)
	go ctx.workspace.SeniorGraphicDesigner.Work(context)
	go ctx.workspace.OperationManager.Work(context)
}
