package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	Mdb *mongo.Collection
	Ctx context.Context
)

func init() {
	Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	// 替换为你自己的 MongoDB 连接 URI
	client, err := mongo.Connect(Ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("连接 MongoDB 失败:", err)
	}

	// 选择数据库和集合
	Mdb = client.Database("testdb").Collection("people")
}
