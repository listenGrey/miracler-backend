package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
)

type person struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func TestMongoSet(t *testing.T) {
	p1 := &person{Name: "Alice", Age: 23}
	p2 := &person{Name: "Tom", Age: 25}
	_, err := Mdb.InsertOne(Ctx, p1)
	if err != nil {
		t.Fatalf("插入失败: %s", err)
	}
	_, err = Mdb.InsertOne(Ctx, p2)
	if err != nil {
		t.Fatalf("插入失败: %s", err)
	}
	t.Logf("插入成功")
}

func TestMongoGet(t *testing.T) {
	cursor, err := Mdb.Find(Ctx, bson.M{})
	if err != nil {
		log.Println("查询失败:", err)
		return
	}
	defer cursor.Close(Ctx)

	fmt.Println("查询结果：")
	for cursor.Next(Ctx) {
		var p person
		if err := cursor.Decode(&p); err != nil {
			t.Fatalf("解码失败:%s", err)
		}
		t.Logf("%v", p)
	}
}

func TestMongoUpdate(t *testing.T) {
	filter := bson.M{"name": "Alice"}
	update := bson.M{"$set": bson.M{"age": 31}}

	res, err := Mdb.UpdateOne(Ctx, filter, update)
	if err != nil {
		t.Fatalf("更新失败: %s", err)
	}

	t.Logf("更新成功: %+v\n", res.ModifiedCount)
}

func TestMongoDel(t *testing.T) {
	filter := bson.M{"name": "Tom"}
	res, err := Mdb.DeleteOne(Ctx, filter)
	if err != nil {
		t.Fatalf("删除失败: %s", err)
	}

	t.Logf("删除成功: %+v\n", res.DeletedCount)
}
