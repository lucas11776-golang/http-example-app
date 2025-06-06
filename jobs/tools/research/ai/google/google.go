package google

import "context"

type GoogleAI struct {
	context context.Context
}

// Comment
func NewGoogleAI(context context.Context) *GoogleAI {
	return &GoogleAI{
		context: context,
	}
}

// Comment
func (ctx *GoogleAI) Text(prompt string) (string, error) {
	return "", nil
}

// Comment
func (ctx *GoogleAI) Audio(prompt string) ([]byte, error) {
	return nil, nil
}

// Comment
func (ctx *GoogleAI) Video(prompt string) ([]byte, error) {
	return nil, nil
}
