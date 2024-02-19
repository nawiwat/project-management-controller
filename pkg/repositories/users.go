package repositories

import (
	"app-controller/pkg/model"
	"context"
)

type  UsersRepository interface {
	Create(ctx context.Context, in model.User) (error)
	Update(ctx context.Context, in model.User) (error)
	UpdateProfile(ctx context.Context, in model.ProfileAttachment) (error)
	Query(ctx context.Context) ([]model.User, error)
	QueryInfo(ctx context.Context,username string) (model.User, error)
	QueryByUsername(ctx context.Context,username string) (model.User, error)
	CreateToken(ctx context.Context, in model.UserToken) (error)
	QueryToken(ctx context.Context,username string) (model.UserToken, error)
}
