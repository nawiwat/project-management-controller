package repositories

import (
	"app-controller/pkg/model"
	"context"
)

type  UsersRepository interface {
	Create(ctx context.Context, in model.User) (error)
	Query(ctx context.Context) ([]model.User, error)
}
