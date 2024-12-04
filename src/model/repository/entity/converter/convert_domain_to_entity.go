package converter

import (
	"github.com/matheusfrteixeira/crud-go/src/model"
	"github.com/matheusfrteixeira/crud-go/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassord(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
