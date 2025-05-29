package workspace

import "server/jobs/office/employees"

type Workspace struct {
	OperationManager      employees.OperationManager
	SeniorAnalyst         employees.SeniorAnalyst
	SeniorGraphicDesigner employees.SeniorGraphicDesigner
	JuniorAnalyst         employees.JuniorAnalyst
}
