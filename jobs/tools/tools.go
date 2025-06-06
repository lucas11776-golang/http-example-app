package tools

import (
	"context"
	"server/env"
	"server/jobs/tools/research"
	"server/utils/prompt"
)

type Tools struct {
	Research *research.Research
	Prompt   *prompt.Prompt
}

// Comment
func NewTools(context context.Context) *Tools {
	return &Tools{
		Research: research.NewResearch(context),
		Prompt:   prompt.NewPrompt(env.Env("PROMPT_PATH"), env.Env("PROMPT_EXTENSION")),
	}
}
