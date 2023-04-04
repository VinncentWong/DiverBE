package util

import "github.com/gin-gonic/gin"

func SendResponse(c *gin.Context, statusCode int, message string, success bool, data any) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"success": success,
		"data":    data,
	})
}
