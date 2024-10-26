package internal

import (
	"snippet-saver/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		// get list of all code snippet
		v1.GET("/snippets/:id",handlers.GetAllSnippet)
		// create code snippet
		v1.POST("/snippets/:id",handlers.CreateSnippet)
		// get snippet by tags
		v1.GET("/snippets/tags/:tag")
		// get snippet by language
		v1.GET("/snippets/language/:lang")
	}
}