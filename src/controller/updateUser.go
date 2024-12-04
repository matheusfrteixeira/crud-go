package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusfrteixeira/crud-go/src/configuration/logger"
	"github.com/matheusfrteixeira/crud-go/src/configuration/rest_err"
	"github.com/matheusfrteixeira/crud-go/src/configuration/validation"
	"github.com/matheusfrteixeira/crud-go/src/controller/model/request"
	"github.com/matheusfrteixeira/crud-go/src/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser controller",
		zap.String("journey", "updateUser"),
	)
	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"),
		)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Param("userId")
	if _, err := bson.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("UserId is not a valid id")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("journey", "updateUser"),
	)

	c.Status(http.StatusOK)
}
