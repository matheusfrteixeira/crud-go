package service

import (
	"github.com/matheusfrteixeira/crud-go/src/configuration/logger"
	"github.com/matheusfrteixeira/crud-go/src/configuration/rest_err"
	"github.com/matheusfrteixeira/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID services", zap.String("journey", "findUSerByID"))

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services", zap.String("journey", "findUserByEmail"))
	return ud.userRepository.FindUserByEmail(email)
}
