package resp

import (
	"alpha/pkg/resp/code"
	"github.com/gin-gonic/gin"
)

type commonRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, commonRes{
		Code:    code.Success,
		Message: code.GetMsg(code.Success),
		Data:    data,
	})
}

func Fail(c *gin.Context, codeVal int) {
	c.JSON(200, commonRes{
		Code:    codeVal,
		Message: code.GetMsg(codeVal),
	})
}

func FailWithMsg(c *gin.Context, codeVal int, msg string) {
	c.JSON(200, commonRes{
		Code:    codeVal,
		Message: msg,
	})
}
