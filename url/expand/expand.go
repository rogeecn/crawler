// Package expand 根据当前路径和当前的链接相对地址获取全路径地址
package expand

import (
	"fmt"
	"net/url"
	"strings"
)

// ExpandBatch ...
func ExpandBatch(currentURL string, dstPath []string) ([]string, error) {
	var fullURLList []string
	for _, item := range dstPath {
		fullItem, err := Expand(currentURL, item)
		if err != nil {
			continue
		}
		fullURLList = append(fullURLList, fullItem)
	}
	return fullURLList, nil
}

// Expand ...
func Expand(currentURL, dstPath string) (string, error) {
	u, err := url.Parse(currentURL)
	if err != nil {
		return "", err
	}

	// url like (http[s]://xxx.xxx/asdfasf) is the full url
	if strings.HasPrefix(dstPath, "http") {
		return dstPath, nil
	}

	// url like (//asdfba/asdfasdf.html) return currentURL's scheme and dstPath
	if strings.HasPrefix(dstPath, "//") {
		return fmt.Sprintf("%s:%s", u.Scheme, dstPath), nil
	}

	// url like (/asdfba/asdfasdf.html) return currentURL's Host and dstPath
	if strings.HasPrefix(dstPath, "/") {
		return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, dstPath), nil
	}

	pathSplit := strings.Split(u.Path, "/")
	pathSplit[len(pathSplit)-1] = dstPath

	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, strings.Join(pathSplit, "/")), nil
}
