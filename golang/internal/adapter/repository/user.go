package repository

import (
	"context"
	"github.com/madsilver/api-challenge/internal/entity"
	"github.com/madsilver/api-challenge/internal/usecase"
)

type UserRepository struct {
	local []entity.User
}

func NewUserRepository() usecase.UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) StoreUsers(ctx context.Context, users []entity.User) error {
	u.local = users
	return nil
}

func (u *UserRepository) GetUsers(ctx context.Context) ([]entity.User, error) {
	return u.local, nil
}
