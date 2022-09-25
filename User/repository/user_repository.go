package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateUser(ctx context.Context, user UserEntity) error
	UpdateUser(ctx context.Context, user UserEntity) error
	// DeleteUser(ctx context.Context, id uuid.UUID) error
	// GetUser(ctx context.Context, id uuid.UUID) (UserEntity, error)
	// GetUsers(ctx context.Context) ([]UserEntity, error)
}

type UserRepository struct {
	MongoDB *mongo.Database
}

func (repo *UserRepository) CreateUser(ctx context.Context, user UserEntity) error {
	_, err := repo.MongoDB.Collection("users").InsertOne(ctx, bson.D{
		bson.E{Key: "name", Value: user.Name},
	})
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) UpdateUser(ctx context.Context, user UserEntity) error {
	_, err := repo.MongoDB.Collection("users").UpdateByID(ctx, user.ID, bson.D{
		{Key: "$set", Value: bson.D{{Key: "name", Value: user.Name}}},
	}, nil)
	if err != nil {
		return err
	}

	return nil
}
