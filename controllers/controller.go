package controllers

import (
	"alpha/logic"
	"alpha/models"
	"alpha/pkg/redis"
	"github.com/gin-gonic/gin"
	"net/http"
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
	// 模拟获取到用户ID
	uid := 1

	// 模拟获取用户的所有事件id
	eventsID := models.MockEventsID

	events, err := logic.GetEvents(redis.Ctx, uid, eventsID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "事件未找到"})
	}
	c.JSON(http.StatusOK, &events)
}

func UpdateItem(c *gin.Context) {
	eventIndex := c.Param("eventId")
	itemIndex := c.Param("itemIndex")

	// 模拟获取到用户ID
	uid := "1"

	var body struct {
		Checked bool `json:"checked"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := logic.UpdateItem(redis.Ctx, uid, eventIndex, itemIndex, body.Checked)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "更新失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "更新成功"})
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
