package models

type DatabaseWord struct {
	Word     string `json:"word"`
	Meanings string `json:"meanings"`
}

type JSONWord struct {
	Word     string   `json:"word"`
	Meanings []string `json:"meanings"`
}
