package repository

import (
	"context"
	"os"

	"github.com/matheusfrteixeira/crud-go/src/configuration/logger"
	"github.com/matheusfrteixeira/crud-go/src/configuration/rest_err"
	"github.com/matheusfrteixeira/crud-go/src/model"
	"github.com/matheusfrteixeira/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUser repository", zap.String("journey", "updateUser"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseCOnnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)
	userIdHex, _ := bson.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil

}
