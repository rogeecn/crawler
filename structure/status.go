package structure

// StatusMatch ...
type StatusMatch struct {
	Working  string `json:"working"`
	Complete string `json:"complete"`
	Stopped  string `json:"stopped"`
}
