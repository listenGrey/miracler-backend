package models

import (
	"fmt"
	"strconv"
)

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

var MockDataTodayEvents = []Event{
	{
		ID:    1,
		Title: "语言",
		Color: "green",
		Items: []Item{
			{
				ID:      1,
				Text:    "背单词",
				Checked: false,
			},
			{
				ID:      2,
				Text:    "听力练习",
				Checked: true,
			},
		},
	},
	{
		ID:    2,
		Title: "健身",
		Color: "blue",
		Items: []Item{
			{
				ID:      1,
				Text:    "深蹲10次",
				Checked: false,
			},
			{
				ID:      2,
				Text:    "跳绳5分钟",
				Checked: false,
			},
		},
	},
}

func PrintMock() {
	fmt.Println(MockDataTodayEvents)
}

func MockUpdateItem(eventID, itemID int, status bool) {
	for i := 0; i < len(MockDataTodayEvents); i++ {
		if MockDataTodayEvents[i].ID == eventID {
			for j := 0; j < len(MockDataTodayEvents[i].Items); j++ {
				if MockDataTodayEvents[i].Items[j].ID == itemID {
					MockDataTodayEvents[i].Items[j].Checked = status
				}
			}
		}
	}
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
