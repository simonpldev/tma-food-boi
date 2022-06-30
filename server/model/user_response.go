package model

import "time"

type UserResponse struct {
	ID        primitive.ObjectID `json:"_id"`
	Name      string             `json:"name"`
	CreatedAt time.Time          `json:"created_at"`
}
