package compress

import (
	"net/url"
)

//Compress ...
func Compress(link string) string {
	// http[s]://www.baidu.com/abc/aas/ccc.html

	return link
}

// Path ...
func Path(link string) string {
	url, err := url.Parse(link)
	if err != nil {
		return link
	}

	return url.Path
}
