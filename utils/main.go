package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "65.49.194.5:6379", // Redis 服务器地址
		Password: "tinghui0430",      // Redis 密码
		DB:       0,                  // 使用第 0 个逻辑数据库
		PoolSize: 10,                 // 连接池大小
	})

	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("❌ Redis 连接失败:", err)
		return
	}
	fmt.Println("✅ Redis 连接成功:", pong)

	// 示例：设置键值对
	err = rdb.Set(ctx, "name", "Alice", 10*time.Minute).Err()
	if err != nil {
		panic(err)
	}

	// 示例：获取值
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name =", val)
}
