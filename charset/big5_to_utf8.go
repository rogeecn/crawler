package charset

import (
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

// Big5ToUtf8 convert big5 content to utf8
func Big5ToUtf8(raw string) (string, error) {
	var charDecoding transform.Transformer
	charDecoding = traditionalchinese.Big5.NewDecoder()

	result, _, err := transform.String(charDecoding, raw)
	if err != nil {
		return raw, err
	}
	return result, nil
}
