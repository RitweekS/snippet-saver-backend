package services

import (
	"snippet-saver/internal/dto/request"
	"snippet-saver/internal/dto/response"
	"snippet-saver/internal/repositories"
)

var AuthInstance Auth = &AuthImpl{}

type Auth interface {
	SignIn(requestBody request.SingInRequest)(*response.SignInResponse, error)
}
type AuthImpl struct{}


func (n *AuthImpl) SignIn(requestBody request.SingInRequest)(res *response.SignInResponse,error error){


	id,err := repositories.UserRepositoriesInstance.SingIn(requestBody)
	
	if err != nil{
		return nil,err
	}

	return &response.SignInResponse{
		Id: id,
	},nil

}