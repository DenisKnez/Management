package repository

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (t *MockUserRepository) CreateUser(ctx context.Context, user UserEntity) error {
	args := t.Called(user)
	return args.Error(0)
}


func (t *MockUserRepository) UpdateUser(ctx context.Context, user UserEntity) error {
	args := t.Called(user)
	return args.Error(0)
}

