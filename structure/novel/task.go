package novel

// Task ...
type Task struct {
	Website      []string `json:"website"`
	NovelInfo    []string `json:"novel_info"`
	NovelChapter []string `json:"novel_chapter"`
	ChapterList  []string `json:"chapter_list"`
}
