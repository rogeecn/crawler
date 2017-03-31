package novel

import (
	"../../structure"
)

// ChapterLinkList ...
type ChapterLinkList struct {
	EntranceURL         string                  `json:"entrance_url"`
	Charset             string                  `json:"charset"`
	LastID              string                  `json:"last_id"`
	ChapterLinkSelector structure.QueryRule     `json:"chapter_link_selector"`
	Rule                structure.LinkFetchRule `json:"rule"`
}
