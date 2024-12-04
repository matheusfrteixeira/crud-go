package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusfrteixeira/crud-go/src/model/service"
)

func NewUserControllerInterface(service service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: service,
	}
}

type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
