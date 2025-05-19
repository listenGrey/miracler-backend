package controllers

import (
	"alpha/models"
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

	eID := models.StringToInt(eventID)
	iID := models.StringToInt(itemIndex)
	models.MockUpdateItem(eID, iID, body.Checked)
	models.PrintMock()

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
