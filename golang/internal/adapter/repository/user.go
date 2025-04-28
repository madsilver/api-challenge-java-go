package repository

import (
	"context"
	"github.com/madsilver/api-challenge/internal/entity"
)

type UserRepository interface {
	StoreUsers(ctx context.Context, users []entity.User) error
	GetUsers(ctx context.Context) ([]entity.User, error)
}
