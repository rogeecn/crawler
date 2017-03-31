package content

import (
	"github.com/rogeecn/crawler/structure"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
)

func TestQueryPath(t *testing.T) {
	var err error
	var result string
	var pattern string
	content := getContent()

	gQuery, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		t.Error(err)
		return
	}

	rule := structure.QueryRule{
		Selector: "meta[property='og:novel:author']",
		Index:    0,
		Type:     "attr",
		Attr:     "content",
	}
	result, err = MatchByQueryPath(gQuery, rule)
	t.Log(pattern, "match: ", result, err)
}
