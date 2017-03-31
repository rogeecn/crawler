package link

import "fmt"

const (
	// OrderAsc ...
	OrderAsc = iota
	// OrderDesc ...
	OrderDesc
)

// NewLinks 链接列表中差异化的最新链接地址
func NewLinks(allLinks []string, lastID string, order int) ([]string, error) {

	// 如果是倒序的需要转成正序列表
	if order == OrderDesc {
		allLinks = reverseLinks(allLinks)
	}

	keyIndex := -1
	for index, linkItem := range allLinks {
		if linkItem == lastID {
			keyIndex = index
			break
		}

	}

	// 如果是-1返回所有LINKS
	if keyIndex == -1 {
		return allLinks, nil
	}

	if keyIndex == len(allLinks)-1 {
		return nil, fmt.Errorf("no new links matched")
	}
	return allLinks[keyIndex+1: len(allLinks)], nil
}

func reverseLinks(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverseLinks(input[1:]), input[0])
}
