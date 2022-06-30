package model

import "time"

type UserCreatePayload struct {
	Name string `json:"name" valid:"required, stringlength(1|50)"`
}

type UserUpdatePayload struct {
	Name      string    `json:"name" valid:"required, stringlength(1|50)"`
	CreatedAt time.Time `json:"created_at"`
}
