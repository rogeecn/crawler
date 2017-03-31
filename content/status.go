package content

import "github.com/rogeecn/crawler/structure"
import "strings"

const (
	// StatusUnknown 未知
	StatusUnknown = "-1"
	// StatusWorking 进行中
	StatusWorking = "0"
	// StatusComplete 完结
	StatusComplete = "1"
	// StatusStopped 停更
	StatusStopped = "2"
)

// CheckStatus ...
func CheckStatus(rawString string, rule structure.StatusMatch) string {
	if strings.Contains(rawString, rule.Working) {
		return StatusWorking
	}

	if strings.Contains(rawString, rule.Complete) {
		return StatusComplete
	}

	if strings.Contains(rawString, rule.Stopped) {
		return StatusStopped
	}

	return StatusUnknown
}
