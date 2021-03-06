package novel

import (
	"github.com/rogeecn/crawler/structure"
)

// BookInfo ...
type BookInfo struct {
	EntranceURL string                         `json:"entrance_url"`
	Charset     string                         `json:"charset"`
	Columns     map[string]structure.QueryRule `json:"columns"`
	Status      structure.StatusMatch          `json:"status"`
}
