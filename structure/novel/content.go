package novel

import (
	"../../structure"
)

// Content ...
type Content struct {
	EntranceURL string              `json:"entrance_url"`
	Charset     string              `json:"charset"`
	Title       structure.QueryRule `json:"title"`
	Content     structure.QueryRule `json:"content"`
}
