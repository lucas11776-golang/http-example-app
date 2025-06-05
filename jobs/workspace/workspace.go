package workspace

import (
	"server/env"
	"server/jobs/office/employees"
	"server/utils/prompt"
)

type Workspace struct {
	OperationManager      employees.OperationManager
	SeniorAnalyst         employees.SeniorAnalyst
	SeniorGraphicDesigner employees.SeniorGraphicDesigner
	JuniorAnalyst         employees.JuniorAnalyst
	Prompt                *prompt.Prompt
}

// Comment
func NewWorkspace() *Workspace {
	return &Workspace{
		Prompt: prompt.NewPrompt(env.Env("PROMPT_PATH"), env.Env("PROMPT_EXTENSION")),
	}
}
