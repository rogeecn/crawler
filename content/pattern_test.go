package content

import (
	"testing"
)

func TestPatternMatch(t *testing.T) {
	var err error
	var pattern string
	content := getContent()

	pattern = `/book/\d+/`
	result, err := MatchByPattern(pattern, content)
	t.Log(pattern, "match: ", result, err)

	results, err := MatchAllByPattern(pattern, content)
	t.Log(pattern, "match: ", results, err)
}
