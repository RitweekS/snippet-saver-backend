package repositories

import (
	"errors"
	"snippet-saver/internal/database"
	"snippet-saver/internal/dto/request"
	"snippet-saver/internal/models"

	"gorm.io/gorm"
)

var UserRepositoriesInstance UserRepositories = &UserRepositoriesImpl{}

type UserRepositories interface {
	SingIn(requestBody request.SingInRequest)(id int,error error)
}

type UserRepositoriesImpl struct{}

func (n *UserRepositoriesImpl) SingIn(requestBody request.SingInRequest)(id int,error error){
	var user models.User

	if err := database.DB.Table("users").Where("email = ?",requestBody.Email).First(&user).Error; err !=nil{

		if errors.Is(err, gorm.ErrRecordNotFound) {
			newUser := models.User{
				Name:  requestBody.Name,
				Email: requestBody.Email,
				Image: requestBody.Image,
			}

			if err := database.DB.Table("users").Create(&newUser).Error;err != nil{{
				return 0, err
			}}

			return int(newUser.ID), nil

		}
		return 0,err
	}

	return int(user.ID),nil
}