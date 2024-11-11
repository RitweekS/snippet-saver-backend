package handlers

import (
	"fmt"
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
		utilities.ApiResponse(http.StatusBadRequest, bindJsonErr.Error(), nil, c)
		return
	}

	userId, userIdErr := utilities.GetUserID(c)

	if userIdErr != nil {
		log.Println(http.StatusBadRequest, userIdErr.Error())
		utilities.ApiResponse(http.StatusBadRequest, userIdErr.Error(), nil, c)
		return
	}

	_, createSnippetsErr := services.SnippetInstance.CreateSnippet(userId, requestBody)

	if createSnippetsErr != nil {
		log.Println(http.StatusInternalServerError, createSnippetsErr.Error())
		utilities.ApiResponse(http.StatusInternalServerError, createSnippetsErr.Error(), nil, c)
		return
	}

	utilities.ApiResponse(http.StatusOK, "Snippet Created Successfully", true, c)
}

func GetAllSnippet(c *gin.Context) {
	userId, userIdErr := utilities.GetUserID(c)
	fmt.Println("serr", userId)
	if userIdErr != nil {
		log.Println(http.StatusBadRequest, userIdErr)
		utilities.ApiResponse(http.StatusBadRequest, userIdErr.Error(), response.GetSnippetResponse{}, c)
	}

	snippets, snippetsErr := services.SnippetInstance.GetAllSnippet(userId)

	if snippetsErr != nil {
		log.Println(http.StatusInternalServerError, snippetsErr.Error())
		utilities.ApiResponse(http.StatusInternalServerError, snippetsErr.Error(), nil, c)
		return
	}

	utilities.ApiResponse(http.StatusOK, "Operation executed Successfully", snippets, c)
}

func GetSnippetByID(c *gin.Context) {
	userId, userIdErr := utilities.GetUserID(c)

	if userIdErr != nil {
		log.Println(http.StatusInternalServerError, userIdErr.Error())
		utilities.ApiResponse(http.StatusInternalServerError, userIdErr.Error(), nil, c)
		return
	}

	snippetId, snippetIdErr := strconv.Atoi(c.Param("snippet_id"))

	if snippetIdErr != nil {
		log.Println(http.StatusInternalServerError, snippetIdErr.Error())
		utilities.ApiResponse(http.StatusInternalServerError, snippetIdErr.Error(), nil, c)
		return
	}

	snippet, snippetErr := services.SnippetInstance.GetSnippetById(userId, snippetId)

	if snippetErr != nil {
		log.Println(http.StatusInternalServerError, snippetIdErr)
		utilities.ApiResponse(http.StatusInternalServerError, snippetIdErr.Error(), nil, c)
		return
	}

	utilities.ApiResponse(http.StatusOK, "Operation executed Successfully", snippet, c)
}
func UpdateSnippetById(c *gin.Context) {
	var requestBody request.CreateSnippetRequest

	parseBodyErr := c.BindJSON(&requestBody)
	if parseBodyErr != nil {
		log.Println("unable to parse", parseBodyErr)
		utilities.ApiResponse(http.StatusBadRequest, parseBodyErr.Error(), nil, c)
		return
	}
	userId, userIdErr := utilities.GetUserID(c)

	if userIdErr != nil {
		log.Println(http.StatusInternalServerError, userIdErr.Error())
		utilities.ApiResponse(http.StatusInternalServerError, userIdErr.Error(), nil, c)
		return
	}

	snippetId, snippetIdErr := strconv.Atoi(c.Param("snippet_id"))

	if snippetIdErr != nil {
		log.Println(http.StatusInternalServerError, snippetIdErr.Error())
		utilities.ApiResponse(http.StatusInternalServerError, snippetIdErr.Error(), nil, c)
		return
	}

	_, snippetErr := services.SnippetInstance.UpdateSnippet(userId, snippetId, requestBody)

	if snippetErr != nil {
		log.Println(http.StatusInternalServerError, snippetIdErr)
		utilities.ApiResponse(http.StatusInternalServerError, snippetErr.Error(), nil, c)
		return
	}

	utilities.ApiResponse(http.StatusOK, "Operation executed Successfully", true, c)

}
func DeleteSnippetByID(c *gin.Context) {
	userId, userIdErr := utilities.GetUserID(c)

	if userIdErr != nil {
		log.Println(http.StatusInternalServerError, userIdErr.Error())
		utilities.ApiResponse(http.StatusInternalServerError, userIdErr.Error(), nil, c)
		return
	}

	snippetId, snippetIdErr := strconv.Atoi(c.Param("snippet_id"))

	if snippetIdErr != nil {
		log.Println(http.StatusInternalServerError, snippetIdErr.Error())
		utilities.ApiResponse(http.StatusInternalServerError, snippetIdErr.Error(), nil, c)
		return
	}

	snippet, snippetErr := services.SnippetInstance.DeleteSnippetById(userId, snippetId)

	if snippetErr != nil {
		log.Println(http.StatusInternalServerError, snippetIdErr)
		utilities.ApiResponse(http.StatusInternalServerError, snippetIdErr.Error(), nil, c)
		return
	}

	utilities.ApiResponse(http.StatusOK, "Operation executed Successfully", snippet, c)
}
