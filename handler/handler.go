package handler

import (
	"alpha/logic"
	"alpha/models"
	"alpha/pkg/redis"
	"alpha/pkg/resp/code"
	"alpha/pkg/resp/resp"
	"github.com/gin-gonic/gin"
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
		resp.FailWithMsg(c, code.NotFound, "事件未找到")
		return
	}

	resp.Success(c, &events)
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
		resp.Fail(c, code.BadRequest)
		return
	}

	err := logic.UpdateItem(redis.Ctx, uid, eventIndex, itemIndex, body.Checked)
	if err != nil {
		resp.FailWithMsg(c, code.BadRequest, "更新失败")
		return
	}

	resp.Success(c, nil)
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
