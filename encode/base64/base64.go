package encode

import "encoding/base64"

// Encode ...
func Encode(rawString string) string {
	return base64.StdEncoding.EncodeToString([]byte(rawString))

}

//Decode ...
func Decode(rawString string) (string, error) {
	byteData, err := base64.StdEncoding.DecodeString(rawString)
	if err != nil {
		return "", err
	}
	return string(byteData), nil
}
