package main

import (
	"github.com/matheusfrteixeira/crud-go/src/controller"
	"github.com/matheusfrteixeira/crud-go/src/model/repository"
	"github.com/matheusfrteixeira/crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InitDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	//Init Dependencies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDoaminService(repo)
	return controller.NewUserControllerInterface(service)
}
