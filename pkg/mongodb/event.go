package mongodb

import (
	"alpha/models"
	"errors"
)

func GetEvent(uid, eventId int) (*models.Event, error) {
	if eventId == MockData1.ID {
		return &MockData1, nil
	} else if eventId == MockData2.ID {
		return &MockData2, nil
	} else {
		return nil, errors.New("event not found")
	}
}

func UpdateItem(uid, eventId, itemId string, checked bool) error {

}

var (
	MockData1 = models.Event{
		ID:    3,
		Title: "清洁",
		Color: "red",
		Items: []models.Item{
			{
				ID:      1,
				Text:    "洗衣服",
				Checked: true,
			},
			{
				ID:      2,
				Text:    "洗澡",
				Checked: true,
			},
		},
	}
	MockData2 = models.Event{
		ID:    4,
		Title: "恋爱",
		Color: "pink",
		Items: []models.Item{
			{
				ID:      1,
				Text:    "约会",
				Checked: true,
			},
			{
				ID:      2,
				Text:    "做爱",
				Checked: false,
			},
		},
	}
)
