package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	Ctx context.Context
	Rdb *redis.Client
)

func init() {
	Ctx = context.Background()
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "65.49.194.5:6379", // Redis 服务器地址
		Password: "tinghui0430",      // Redis 密码
		DB:       0,                  // 使用第 0 个逻辑数据库
		PoolSize: 10,                 // 连接池大小
	})
}
