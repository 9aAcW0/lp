package models

type Argument struct {
	Content  string   `json:"content"`
	Contents []string `json:"contents"`
	Pointer  []string `json:"pointer"`
}
