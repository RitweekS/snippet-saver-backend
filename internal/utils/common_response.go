package utilities

import (
	"github.com/gin-gonic/gin"
)

func ApiResponse(responseCode int, responseMessage string, responseData interface{}, c *gin.Context) {

	c.JSON(responseCode, gin.H{
		"response_code":    responseCode,
		"response_message": responseMessage,
		"response_data":    responseData,
	})
}
