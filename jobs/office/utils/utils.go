package utils

import "regexp"

var (
	PAPERWORK_RESULT_REGEX = regexp.MustCompile(`<result>(?s:(.*?))</result>`)
)

// Comment
func ResultFromPaperword(document string) string {
	result := PAPERWORK_RESULT_REGEX.FindStringSubmatch(document)

	if len(result) == 0 {
		return ""
	}

	return result[1]
}
