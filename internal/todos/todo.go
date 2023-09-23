package todos

type Todo struct {
	Filename    string `json:"filename"`
	PackageName string `json:"packageName"`
	Line        int    `json:"line"`
	Text        string `json:"text"`
}
