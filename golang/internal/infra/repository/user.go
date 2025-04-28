package repository

import (
	"context"
	"github.com/madsilver/api-challenge/internal/adapter/repository"
	"github.com/madsilver/api-challenge/internal/entity"
)

type UserRepository struct {
	local []entity.User
}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) StoreUsers(ctx context.Context, users []entity.User) error {
	u.local = users
	return nil
}

func (u *UserRepository) GetUsers(ctx context.Context) ([]entity.User, error) {
	return u.local, nil
}
