package utils

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, code int, message string, data interface{}) {
	theResponse := gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	}
	c.JSON(code, theResponse)
}

func Failed(c *gin.Context, code int, message string, data interface{}) {
	theResponse := gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	}
	c.JSON(code, theResponse)
}
