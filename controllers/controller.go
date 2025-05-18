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

func Today(c *gin.Context) {
	today := models.MockDataToday
	c.JSON(http.StatusOK, today)
}

func Summary(c *gin.Context) {

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

func Calendar(c *gin.Context) {

}

func ModifyOneDay(c *gin.Context) {

}

func NewEvent(c *gin.Context) {

}

func Overview(c *gin.Context) {

}

func Profile(c *gin.Context) {

}

func ModifyEvent(c *gin.Context) {

}

func DeleteEvent(c *gin.Context) {

}
