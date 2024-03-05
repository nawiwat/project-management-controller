package repositories

import (
	"app-controller/pkg/model"
	"context"
)

type  TasksRepository interface {
	Create(ctx context.Context, in model.Task ) ([]model.Task , error)
	Query(ctx context.Context, id uint64 ) ([]model.Task , error)
	Update(ctx context.Context, in []model.Task ) ([]model.Task , error)
	Delete(ctx context.Context, in uint64 ) (error)
	QueryByUserId(ctx context.Context, id uint64 ) ([]model.Task , error)
}