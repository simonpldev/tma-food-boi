package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"created_at"`
}

type UserCreateBSON struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"created_at"`
}

type UserUpdateBSON struct {
	Name      string    `bson:"name"`
	CreatedAt time.Time `bson:"created_at"`
}

func (payload UserCreatePayload) ConvertToCreateBSON() (bson UserCreateBSON) {
	bson = UserCreateBSON{
		ID:        primitive.NewObjectID(),
		Name:      payload.Name,
		CreatedAt: time.Now(),
	}
	return
}

func (payload UserUpdatePayload) ConvertToUpdateBSON() (bson UserUpdateBSON) {
	bson = UserUpdateBSON{
		Name:      payload.Name,
		CreatedAt: payload.CreatedAt,
	}
	return
}
