package content

import "strings"

// Replace ...
func Replace(rawString, replaceChars string) string {
	replaceCharsArr := strings.Split(replaceChars, ",")

	var replaceParams []string
	for _, char := range replaceCharsArr {
		replacePair := strings.Split(char, "=>")
		if len(replacePair) == 2 {
			replaceParams = append(replaceParams, replacePair...)
			continue
		}
		replaceParams = append(replaceParams, char, "")
	}

	stringReplacer := strings.NewReplacer(replaceParams...)
	return stringReplacer.Replace(rawString)
}
