package redis

import (
	"alpha/models"
	"encoding/json"
	"testing"
	"time"
)

func TestRedisPing(t *testing.T) {
	res := Rdb.Ping(Ctx)
	if res.Err() != nil {
		t.Fatalf("conn error: %v", res.Err())
	}
	t.Logf("%s", res.Val())
}

func TestRedisGetSet(t *testing.T) {
	err := Rdb.Set(Ctx, "id", "user name", 1*time.Minute).Err()
	if err != nil {
		t.Fatalf("set value error %s", err)
	}

	val, err := Rdb.Get(Ctx, "id").Result()
	if err != nil {
		t.Fatalf("get value error %s", err)
	}
	t.Logf("%s", val)
}

func TestRedisSetTestData(t *testing.T) {
	demo := []models.Event{
		{
			ID:    1,
			Title: "语言",
			Color: "green",
			Items: []models.Item{
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
			Items: []models.Item{
				{
					ID:      1,
					Text:    "深蹲10次",
					Checked: true,
				},
				{
					ID:      2,
					Text:    "跳绳5分钟",
					Checked: false,
				},
			},
		},
	}

	val, err := json.Marshal(demo)
	if err != nil {
		t.Fatalf("convert to json error")
	}

	err = Rdb.Set(Ctx, "event:1", val, 0).Err()
	if err != nil {
		t.Fatalf("set value error %s", err)
	}

	t.Logf("set value success")
}

func TestRedisGetTestData(t *testing.T) {
	val, err := Rdb.Get(Ctx, "event:1").Result()
	if err != nil {
		t.Fatalf("get value error %s", err)
	}
	t.Logf("%v", val)
}
