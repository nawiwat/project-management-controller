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
	CreateNotification(ctx context.Context, in model.Notification) (error)
	GetNotification(ctx context.Context, in uint64) (model.Notification , error)
	DeleteNotification(ctx context.Context, in model.Notification) (error)
	UpdateNotification(ctx context.Context, in model.User , cur_task []model.Task) (error)
}
