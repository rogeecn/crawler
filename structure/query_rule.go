package structure

// QueryRule  如果type=attr则使用attr标签
// 如果存在 pattern 则忽略其它匹配方式
// Type [text/html/attr]
type QueryRule struct {
	Selector string `json:"selector"`
	Index    int    `json:"index"`
	Type     string `json:"type"`
	Attr     string `json:"attr"`
	Pattern  string `json:"pattern"`
	Replace  string `json:"replace"`
}
