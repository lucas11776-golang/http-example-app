package producer

import (
	"context"
	"regexp"
	"server/workers/agents/studio/producer"
)

var (
	RESULT_REGEX = regexp.MustCompile(`<result>(.*?)</result>`)
)

type Reporter struct {
}

type ViewScript struct {
}

// Reporter
func (ctx *Reporter) Action(context context.Context, script *producer.NewsScript) (*ViewScript, error) {
	// Do some reporting...
	return nil, nil
}
