package novel

import (
	"github.com/rogeecn/crawler/structure"
)

// NewBook 新书发现数据结构
type NewBook struct {
	EntranceURL string                  `json:"entrance_url"`
	Charset     string                  `json:"charset"`
	Rule        structure.LinkFetchRule `json:"rule"`
}
