package content

import (
	"../structure"
	"github.com/PuerkitoBio/goquery"
)

// Fetch ...
func Fetch(gQuery *goquery.Document, rule structure.QueryRule) (string, error) {
	var contentData string
	var err error
	if len(rule.Pattern) > 0 {
		html, err := gQuery.Html()
		if nil != err {
			return "", err
		}
		contentData, err = MatchByPattern(rule.Pattern, html)
	} else {
		contentData, err = MatchByQueryPath(gQuery, rule)
	}

	if nil != err {
		return "", err
	}
	return Replace(contentData, rule.Replace), nil
}
