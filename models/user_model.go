package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDB struct {
	Username  string    `json:"username" bson:"username" binding:"required"`
	Email     string    `json:"email" bson:"email" binding:"required"`
	Phone     string    `json:"phone" bson:"phone" binding:"required"`
	Address   string    `json:"address" bson:"address" binding:"required"`
	City      string    `json:"city" bson:"city" binding:"required"`
	CreateAt  time.Time `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UserResponse struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address   string             `json:"address,omitempty" bson:"address,omitempty"`
	City      string             `json:"city,omitempty" bson:"city,omitempty"`
	CreateAt  time.Time          `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UserUpdate struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address   string             `json:"address,omitempty" bson:"address,omitempty"`
	CreateAt  time.Time          `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
