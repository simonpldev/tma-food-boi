package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tma-food-boi/dao"
	"tma-food-boi/model"
)

func UserCreate(payload model.UserCreatePayload) (userID primitive.ObjectID, err error) {
	var (
		userBSON = payload.ConvertToCreateBSON()
	)

	err = dao.UserCreate(userBSON)
	userID = userBSON.ID
	return
}

func UserList() (userListRes []model.UserResponse, err error) {
	users, err := dao.UserList()
	if err != nil {
		return
	}

	for _, user := range users {
		userRes := UserConvertToResponse(user)
		userListRes = append(userListRes, userRes)
	}

	return
}

func UserConvertToResponse(u model.User) (res model.UserResponse) {
	res = model.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
	}
	return
}
