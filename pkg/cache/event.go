package cache

import (
	"alpha/models"
	"fmt"
	"github.com/redis/go-redis/v9"

	rdb "alpha/pkg/redis"
	"context"
	"encoding/json"
)

func GetEvent(ctx context.Context, uid int, eventId int) (*models.Event, error) {
	key := fmt.Sprintf("uid:%d:event:%d", uid, eventId)
	val, err := rdb.Rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		// 未命中缓存
		return nil, err
	} else if err != nil {
		return nil, err
	}

	var event models.Event
	if err := json.Unmarshal([]byte(val), &event); err != nil {
		return nil, err
	}
	return &event, nil
}

func SetEvent(ctx context.Context, uid int, event *models.Event) error {
	key := fmt.Sprintf("uid:%d:event:%d", uid, event.ID)
	b, err := json.Marshal(&event)
	if err != nil {
		return err
	}
	return rdb.Rdb.Set(ctx, key, b, 0).Err()
}

func DelEvent(ctx context.Context, uid, eventId interface{}) error {
	key := fmt.Sprintf("uid:%s:event:%s", uid, eventId)
	return rdb.Rdb.Del(ctx, key).Err()
}
