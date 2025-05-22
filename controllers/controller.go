package controllers

import (
	"alpha/models"
	"alpha/pkg/redis"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {

}

func Login(c *gin.Context) {

}

// Today画面的函数

func Today(c *gin.Context) {

}

func Summary(c *gin.Context) {

}

func Layup(c *gin.Context) {

}

func GetEvents(c *gin.Context) {
	c.JSON(http.StatusOK, models.MockDataTodayEvents)
}

func UpdateItem(c *gin.Context) {
	eventID := c.Param("eventId")
	itemIndex := c.Param("itemIndex")

	var body struct {
		Checked bool `json:"checked"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 在此保存修改:数据库更新，缓存更新
	key := "event:" + eventID
	val, err := redis.Rdb.Get(redis.Ctx, key).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "事件未找到"})
		return
	}

	var event models.Event
	if err := json.Unmarshal([]byte(val), &event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据解析失败"})
		return
	}

	itemId, _ := strconv.Atoi(itemIndex)
	if itemId < 0 || itemId >= len(event.Items) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "事项索引无效"})
		return
	}

	event.Items[itemId].Checked = body.Checked

	updateBytes, _ := json.Marshal(event)
	redis.Rdb.Set(redis.Ctx, key, updateBytes, 0)

	c.JSON(http.StatusOK, gin.H{
		"message":    "更新成功",
		"eventId":    eventID,
		"itemIndex":  itemIndex,
		"newChecked": body.Checked,
	})
}

// Calendar画面的函数

func Calendar(c *gin.Context) {

}

func ModifyOneDay(c *gin.Context) {

}

// New画面的函数

func NewEvent(c *gin.Context) {

}

// Overview画面的函数

func Overview(c *gin.Context) {

}

// Profile画面的函数

func Profile(c *gin.Context) {

}

func ModifyEvent(c *gin.Context) {

}

func DeleteEvent(c *gin.Context) {

}
