package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID   primitive.ObjectID
	Name string
}
