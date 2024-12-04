package entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserEntity struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Email    string        `bson:"_email"`
	Password string        `bson:"_password"`
	Name     string        `bson:"_name"`
	Age      int8          `bson:"_age"`
}
