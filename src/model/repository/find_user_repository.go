package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/matheusfrteixeira/crud-go/src/configuration/logger"
	"github.com/matheusfrteixeira/crud-go/src/configuration/rest_err"
	"github.com/matheusfrteixeira/crud-go/src/model"
	"github.com/matheusfrteixeira/crud-go/src/model/repository/entity"
	"github.com/matheusfrteixeira/crud-go/src/model/repository/entity/converter"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init finUserByEmail repository", zap.String("journey", "finUserByEmail"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseCOnnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_email", Value: email}}

	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("user with email %s not found", email)

			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init finUserById repository", zap.String("journey", "finUserById"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseCOnnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := bson.ObjectIDFromHex(id) //primitive.ObjectIDFromHex(id)

	// fmt.Println(objectId)

	filter := bson.D{{Key: "_id", Value: objectId}}

	// fmt.Println(filter)

	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("user with ID %s not found", id)

			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by ID"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	return converter.ConvertEntityToDomain(*userEntity), nil
}
