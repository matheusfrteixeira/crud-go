package service

import (
	"github.com/matheusfrteixeira/crud-go/src/configuration/logger"
	"github.com/matheusfrteixeira/crud-go/src/configuration/rest_err"
	"github.com/matheusfrteixeira/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userID string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(userID, userDomain)
	if err != nil {
		return err
	}

	return nil
}
