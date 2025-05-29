package office

import (
	"server/jobs/office/employees/analyst/junior"
	"server/jobs/office/employees/analyst/senior"
	"server/jobs/office/employees/manager"
	"server/jobs/office/workspace"
)

type Office struct {
	workspace *workspace.Workspace
}

func NewOffice() *Office {
	workspace := &workspace.Workspace{}

	workspace.OperationManager = manager.NewOperationManager(workspace)
	workspace.SeniorAnalyst = senior.NewSeniorAnalyst(workspace)
	workspace.SeniorGraphicDesigner = senior.NewSeniorAnalyst(workspace)
	workspace.JuniorAnalyst = junior.NewJuniorAnalyst(workspace)

	return &Office{workspace: workspace}
}

func (ctx *Office) Launch() {
	go ctx.workspace.JuniorAnalyst.Work()
	go ctx.workspace.SeniorAnalyst.Work()
	go ctx.workspace.SeniorGraphicDesigner.Work()
	go ctx.workspace.OperationManager.Work()
}
