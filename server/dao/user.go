package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"tma-food-boi/db"
	"tma-food-boi/model"
)

func UserCreate(user model.UserCreateBSON) (err error) {
	var (
		userCol = db.UserCol()
		c       = context.Background()
	)
	_, err = userCol.InsertOne(c, user)
	return
}

func UserList() (userList []model.User, err error) {
	var (
		userCol = db.UserCol()
		c       = context.Background()
	)

	cur, err := userCol.Aggregate(c, bson.M{})
	if err != nil {
		return
	}

	defer cur.Close(c)
	cur.All(c, &userList)

	return
}
