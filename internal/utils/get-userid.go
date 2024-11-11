package utilities

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) (int, error) {
	userID, exists := c.Get("userID")

	if !exists {
		return 0, fmt.Errorf("user ID not found in context")
	}

	id, ok := userID.(int)
	if !ok {
		return 0, fmt.Errorf("user ID is not in expected format")
	}

	return id, nil
}
