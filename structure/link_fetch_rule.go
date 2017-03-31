package structure

// LinkFetchRule ...
type LinkFetchRule struct {
	Selector string    `json:"selector"`
	Attr     string    `json:"attr"`
	Rule     MatchRule `json:"rule"`
	Order    int       `json:"order"`
	Pattern  string    `json:"pattern"`
}
