package encode

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// MD5 ...
func MD5(rawString string) string {
	hasher := md5.New()
	hasher.Write([]byte(rawString))
	return hex.EncodeToString(hasher.Sum(nil))
}

// MD5File returns the md5 for a given file.
func MD5File(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("Error opening file for MD5 computation: %v (%s)", err, filePath))
	}

	hasher := md5.New()
	io.Copy(hasher, f)

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
