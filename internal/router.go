package internal

import (
	"snippet-saver/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		auth := v1.Group("auth")
		{
			auth.POST("/signin",handlers.SignIn)
		}
		snippets:= v1.Group("/snippets")
		{
			snippets.GET("/:id",handlers.GetAllSnippet)
			// create code snippet
			snippets.POST("/:id",handlers.CreateSnippet)
			// get snippet by tags
			snippets.GET("/tags/:tag")
			// get snippet by language
			snippets.GET("/language/:lang")
		}
	}
	
}