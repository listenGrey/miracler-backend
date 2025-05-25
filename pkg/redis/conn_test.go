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

func TestRedisSet(t *testing.T) {
	err := Rdb.Set(Ctx, "id", "user name", 1*time.Minute).Err()
	if err != nil {
		t.Fatalf("set value error %s", err)
	}
	t.Logf("success")
}

func TestRedisGet(t *testing.T) {
	val, err := Rdb.Get(Ctx, "id").Result()
	if err != nil {
		t.Fatalf("get value error %s", err)
	}
	t.Logf("%s", val)
}

func TestRedisDel(t *testing.T) {
	err := Rdb.Del(Ctx, "id").Err()
	if err != nil {
		t.Fatalf("delete value error %s", err)
	}
	t.Logf("success")
}

func TestRedisSetEventTestData1(t *testing.T) {
	demo := models.Event{
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
	}

	val, err := json.Marshal(demo)
	if err != nil {
		t.Fatalf("convert to json error")
	}

	err = Rdb.Set(Ctx, "uid:1:event:1", val, 0).Err()
	if err != nil {
		t.Fatalf("set value error %s", err)
	}

	t.Logf("set value success")
}

func TestRedisGetEventTestData1(t *testing.T) {
	val, err := Rdb.Get(Ctx, "uid:1:event:1").Result()
	if err != nil {
		t.Fatalf("get value error %s", err)
	}

	var event models.Event
	if err := json.Unmarshal([]byte(val), &event); err != nil {
		t.Fatalf("数据解析失败")
	}
	t.Logf("%v", event)
}

func TestRedisDelEventTestData1(t *testing.T) {
	err := Rdb.Del(Ctx, "uid:1:event:1").Err()
	if err != nil {
		t.Fatalf("delete value error %s", err)
	}
	t.Logf("success")
}

func TestRedisSetEventTestData2(t *testing.T) {
	demo := models.Event{
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
	}
	val, err := json.Marshal(demo)
	if err != nil {
		t.Fatalf("convert to json error")
	}

	err = Rdb.Set(Ctx, "uid:1:event:2", val, 0).Err()
	if err != nil {
		t.Fatalf("set value error %s", err)
	}

	t.Logf("set value success")
}

func TestRedisGetEventTestData2(t *testing.T) {
	val, err := Rdb.Get(Ctx, "uid:1:event:2").Result()
	if err != nil {
		t.Fatalf("get value error %s", err)
	}

	var event models.Event
	if err := json.Unmarshal([]byte(val), &event); err != nil {
		t.Fatalf("数据解析失败")
	}
	t.Logf("%v", event)
}

func TestRedisDelEventTestData2(t *testing.T) {
	err := Rdb.Del(Ctx, "uid:1:event:2").Err()
	if err != nil {
		t.Fatalf("delete value error %s", err)
	}
	t.Logf("success")
}
