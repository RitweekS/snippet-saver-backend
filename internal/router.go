package internal

import (
	"snippet-saver/internal/handlers"
	"snippet-saver/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "working",
		})
	})
	v1 := router.Group("/v1")
	{
		auth := v1.Group("auth")
		{
			auth.POST("/signin", handlers.SignIn)
		}
		snippets := v1.Group("/snippets")
		snippets.Use(middleware.Authentication())
		{
			snippets.GET("/", handlers.GetAllSnippet)
			// create code snippet
			snippets.POST("/", handlers.CreateSnippet)

			//get single snippet
			snippets.GET("//:snippet_id", handlers.GetSnippetByID)
			//edit the snippet
			snippets.PUT("//:snippet_id", handlers.UpdateSnippetById)
			//delete the snippet
			snippets.DELETE("/:snippet_id", handlers.DeleteSnippetByID)
		}
	}

}
