package content

import (
	"fmt"
	"regexp"
)

// MatchByPattern match content from pattern
func MatchByPattern(pattern, fullContent string) (string, error) {
	pattern = fmt.Sprintf("(?sim)%s", pattern)
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	result := reg.FindStringSubmatch(fullContent)
	if len(result) == 0 {
		return "", fmt.Errorf("pattern match no content")
	}

	return result[1], nil
}

// MatchByPattern match content from pattern
func MatchAllByPattern(pattern, fullContent string) ([]string, error) {
	pattern = fmt.Sprintf("(?sim)%s", pattern)
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	results := reg.FindAllStringSubmatch(fullContent, -1)
	if len(results) == 0 {
		return nil, fmt.Errorf("pattern match no content")
	}

	var result []string
	for _, item := range results {
		result = append(result, item[1])
	}
	return result, nil
}
