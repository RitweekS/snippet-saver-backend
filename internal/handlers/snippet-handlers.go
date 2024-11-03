package handlers

import (
	"log"
	"net/http"
	"snippet-saver/internal/dto/request"
	"snippet-saver/internal/dto/response"
	"snippet-saver/internal/services"
	utilities "snippet-saver/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)


func CreateSnippet(c *gin.Context) {
	var requestBody request.CreateSnippetRequest

	bindJsonErr := c.BindJSON(&requestBody)

	if bindJsonErr != nil {
		log.Println(http.StatusBadRequest, bindJsonErr.Error())
		utilities.ApiResponse(http.StatusBadRequest,bindJsonErr.Error(),nil,c)
		return
	}

	snippetIdParam,snippetIdError := strconv.Atoi(c.Param("id"))

	if snippetIdError != nil {
		log.Println(http.StatusBadRequest, snippetIdError.Error())
		utilities.ApiResponse(http.StatusBadRequest,snippetIdError.Error(),nil,c)
		return
	}

	_,createSnippetsErr := services.SnippetInstance.CreateSnippet(snippetIdParam,requestBody);

	if createSnippetsErr!= nil {
		log.Println(http.StatusInternalServerError, createSnippetsErr.Error())
		utilities.ApiResponse(http.StatusInternalServerError,createSnippetsErr.Error(),nil,c)
		return
	}
	
	utilities.ApiResponse(http.StatusOK,"Snippet Created Successfully",true,c)
}

func GetAllSnippet(c *gin.Context) {
	userId,userIdErr := strconv.Atoi(c.Param("id"))

	if userIdErr!=nil{
		log.Println(http.StatusBadRequest,userIdErr)
		utilities.ApiResponse(http.StatusBadRequest,"Operation executed Successfully",response.GetSnippetResponse{},c)
	}

	snippets,snippetsErr := services.SnippetInstance.GetAllSnippet(userId)

	if snippetsErr !=nil{
		log.Println(http.StatusInternalServerError, snippetsErr.Error())
		utilities.ApiResponse(http.StatusInternalServerError,snippetsErr.Error(),nil,c)
		return
	}

	utilities.ApiResponse(http.StatusOK,"Operation executed Successfully",snippets,c)
}