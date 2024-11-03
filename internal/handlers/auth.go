package handlers

import (
	"log"
	"net/http"
	"snippet-saver/internal/dto/request"
	"snippet-saver/internal/services"
	utilities "snippet-saver/internal/utils"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var requestBody request.SingInRequest

	bindJSONErr := c.BindJSON(&requestBody)

	if bindJSONErr != nil {
		log.Println(http.StatusBadRequest, bindJSONErr.Error())
		utilities.ApiResponse(http.StatusBadRequest,bindJSONErr.Error(),nil,c)
		return
	}

	res,err := services.AuthInstance.SignIn(requestBody)

	if err!=nil{
		log.Println(http.StatusInternalServerError, err.Error())
		utilities.ApiResponse(http.StatusBadRequest,err.Error(),nil,c)
		return
	}

	if res==nil{
		utilities.ApiResponse(http.StatusOK,"operation executed successfully",nil,c)
		return
	}

	utilities.ApiResponse(http.StatusOK,"operation executed successfully",res,c)
}