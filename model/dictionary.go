package model

type SearchResults struct {
	Hwi      Hwi      `json:"hwi"`
	Fl       string   `json:"fl"`
	Shortdef []string `json:"shortdef"`
}
type Hwi struct {
	Hw string `json:"hw"`
}
