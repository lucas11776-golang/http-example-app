package manager

import (
	"context"
	"regexp"
)

var (
	RESULT_REGEX = regexp.MustCompile(`<result>(.*?)</result>`)
)

type StudioManager struct {
}

type NewsArticle struct {
}

type Studio struct {
}

// Reporter
func (ctx *StudioManager) PrepareStudio(context context.Context, article *NewsArticle) (*Studio, error) {
	// Setup studio base on article...
	return nil, nil
}
