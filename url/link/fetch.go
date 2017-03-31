package link

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rogeecn/crawler/content"
	"github.com/rogeecn/crawler/structure"
)

// FetchLinks ...
func FetchLinks(gQuery *goquery.Document, rule structure.LinkFetchRule) []string {
	if len(rule.Pattern) > 0 {
		return fetchLinksByPattern(rule.Pattern, gQuery)
	}

	var result []string
	gQuery.Find(rule.Selector).Each(func(idx int, item *goquery.Selection) {
		if itemResult, exist := item.Attr(rule.Attr); exist {
			if urlMatchRule(itemResult, rule.Rule) {
				result = append(result, itemResult)
			}
		}
	})

	var retResult []string
	tmpMapStorage := make(map[string]string)
	for _, item := range result {
		if _, ok := tmpMapStorage[item]; !ok {
			tmpMapStorage[item] = "1"
			retResult = append(retResult, item)
		}
	}
	tmpMapStorage = nil

	return retResult
}

// fetchLinksByPattern ...
func fetchLinksByPattern(pattern string, gQuery *goquery.Document) []string {
	html, _ := gQuery.Html()
	result, err := content.MatchAllByPattern(pattern, html)
	if err != nil {
		return nil
	}
	return result
}

func urlMatchRule(url string, rule structure.MatchRule) bool {
	//包含
	if len(rule.Contain) > 0 {
		if !strings.Contains(url, rule.Contain) {
			return false
		}
	}

	if len(rule.NotContain) > 0 {
		if strings.Contains(url, rule.NotContain) {
			return false
		}
	}

	if len(rule.Start) > 0 {
		if !strings.HasPrefix(url, rule.Start) {
			return false
		}
	}

	if len(rule.End) > 0 {
		if !strings.HasSuffix(url, rule.End) {
			return false
		}
	}

	return true
}
