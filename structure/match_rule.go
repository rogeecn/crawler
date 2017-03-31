package structure

// MatchRule ...
type MatchRule struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	Contain    string `json:"contain"`
	NotContain string `json:"not_contain"`
}
