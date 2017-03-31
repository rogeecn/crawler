package compress

import (
	"fmt"
	"net/url"
)

//Compress ...
func Compress(link string) string {
	// http[s]://www.baidu.com/abc/aas/ccc.html
	url, err := url.Parse(link)
	if err != nil {
		return link
	}

	if len(url.RawQuery) == 0 {
		return url.Path
	}
	return fmt.Sprintf("%s?%s", url.Path, url.RawQuery)
}
