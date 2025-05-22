package models

type Today struct {
	Date   string  `json:"date"`
	Events []Event `json:"events"`
}

type Event struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
	Items []Item `json:"items"`
}

type Item struct {
	// Title  string      `json:"title"`
	// Action interface{} `json:"action"`
	ID      int    `json:"id"`
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
}
