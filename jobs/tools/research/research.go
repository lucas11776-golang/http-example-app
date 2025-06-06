package research

import (
	"context"
	"server/jobs/tools/research/ai"
	"server/jobs/tools/research/ai/google"
)

type Research struct {
	ai ai.AI
}

// Comment
func NewResearch(context context.Context) *Research {
	return &Research{
		ai: google.NewGoogleAI(context),
	}
}

// Comment
func (ctx *Research) Text(prompt string) (string, error) {
	return ctx.ai.Text(prompt)
}

// Comment
func (ctx *Research) Audio(prompt string) ([]byte, error) {
	return ctx.ai.Audio(prompt)
}

// Comment
func (ctx *Research) Video(prompt string) ([]byte, error) {
	return ctx.ai.Video(prompt)
}
