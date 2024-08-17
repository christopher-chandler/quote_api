package body

// The basic structure of a quote
// that is to be read in from the csv file.
type Quote struct {
	Text     string `json:"text"`
	Author   string `json:"author"`
	Category string `json:"category"`
}
