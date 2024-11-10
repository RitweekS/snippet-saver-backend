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
			auth.POST("/signin", handlers.SignIn)
		}
		snippets := v1.Group("/snippets")
		{
			snippets.GET("/:user_id", handlers.GetAllSnippet)
			// create code snippet
			snippets.POST("/:user_id", handlers.CreateSnippet)

			//get single snippet
			snippets.GET("/:user_id/:snippet_id", handlers.GetSnippetByID)
			//edit the snippet
			snippets.PUT("/:user_id/:snippet_id", handlers.UpdateSnippetById)
			//delete the snippet
			snippets.DELETE("/:user_id/:snippet_id", handlers.DeleteSnippetByID)
		}
	}

}
