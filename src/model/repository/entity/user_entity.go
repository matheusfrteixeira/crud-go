package entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserEntity struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Email    string        `bson:"_email,omitempty"`
	Password string        `bson:"_password,omitempty"`
	Name     string        `bson:"_name,omitempty"`
	Age      int8          `bson:"_age,omitempty"`
}
