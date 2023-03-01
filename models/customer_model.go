package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerDB struct {
	Username  string            `json:"username" bson:"username" binding:"required"`
	Dni       string            `json:"dni" bson:"dni" binding:"required"`
	Email     string            `json:"email" bson:"email" binding:"required"`
	Phone     string            `json:"phone" bson:"phone" binding:"required"`
	Address   string            `json:"address" bson:"address" binding:"required"`
	City      string            `json:"city" bson:"city" binding:"required"`
	Car       map[string]string `json:"car" bson:"car" binding:"required"`
	CreateAt  time.Time         `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdatedAt time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CustomerResponse struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Dni       string             `json:"dni,omitempty" bson:"dni,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address   string             `json:"address,omitempty" bson:"address,omitempty"`
	City      string             `json:"city,omitempty" bson:"city,omitempty"`
	Car       map[string]string  `json:"car,omitempty" bson:"car,omitempty"`
	CreateAt  time.Time          `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CustomerUpdate struct {
	Username  string            `json:"username,omitempty" bson:"username,omitempty"`
	Email     string            `json:"email,omitempty" bson:"email,omitempty"`
	Phone     string            `json:"phone,omitempty" bson:"phone,omitempty"`
	Address   string            `json:"address,omitempty" bson:"address,omitempty"`
	Car       map[string]string `json:"car,omitempty" bson:"car,omitempty"`
	CreateAt  time.Time         `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdatedAt time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
