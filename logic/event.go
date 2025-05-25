package logic

import (
	"alpha/models"
	"alpha/pkg/cache"
	"alpha/pkg/mongodb"
	"context"
	"github.com/redis/go-redis/v9"
)

func GetEvents(ctx context.Context, uid int, eventsId []int) (*[]models.Event, error) {
	var events []models.Event
	for _, eventId := range eventsId {
		eventCache, err := cache.GetEvent(ctx, uid, eventId)
		if err == nil {
			// 缓存命中，读缓存
			events = append(events, *eventCache)
		} else {
			// 缓存未命中，查询数据库
			eventDB, err := mongodb.GetEvent(uid, eventId)
			if err != nil {
				continue
			}
			events = append(events, *eventDB)

			// 填充缓存
			err = cache.SetEvent(ctx, uid, eventDB)
			if err != nil {
				return nil, err
			}
		}
	}
	return &events, nil
}

func UpdateItem(ctx context.Context, uid, eventId, itemId string, checked bool) error {
	// 删除缓存
	err := cache.DelEvent(ctx, uid, eventId)
	if err != nil && err != redis.Nil {
		return err
	}

	// 写DB
	err = mongodb.UpdateItem(uid, eventId, itemId, checked)
	if err != nil {
		return err
	}

	return nil
}
