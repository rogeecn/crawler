package charset

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GBKToUTF8 convert gbk content to utf8
func GBKToUTF8(raw string) (string, error) {
	var charDecoding transform.Transformer
	charDecoding = simplifiedchinese.GBK.NewDecoder()

	result, _, err := transform.String(charDecoding, raw)
	if err != nil {
		return raw, err
	}
	return result, nil
}
