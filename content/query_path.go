package content

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rogeecn/crawler/structure"
)

// MatchByQueryPath select content data from query path
func MatchByQueryPath(gQuery *goquery.Document, rule structure.QueryRule) (string, error) {
	selector := gQuery.Find(rule.Selector).Eq(rule.Index)
	switch rule.Type {
	case "text":
		return selector.Text(), nil
	case "html":
		return selector.Html()
	case "attr":
		return selector.AttrOr(rule.Attr, ""), nil
	}
	return "", nil
}
