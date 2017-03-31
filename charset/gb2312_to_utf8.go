package charset

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GB2312ToUTF8 convert gb2312 content to utf8
func GB2312ToUTF8(raw string) (string, error) {
	var charDecoding transform.Transformer
	charDecoding = simplifiedchinese.HZGB2312.NewDecoder()

	result, _, err := transform.String(charDecoding, raw)
	if err != nil {
		return raw, err
	}
	return result, nil
}
