package repositories

import (
	"app-controller/pkg/model"
	"context"
)

type  ProjectsRepository interface {
	Create(ctx context.Context, in model.Project) (model.Project, error)
	Update(ctx context.Context, in model.Project) (error)
	AddMember(ctx context.Context, in model.Membership) (error)
	Query(ctx context.Context) ([]model.Project, error)
	QueryInfo(ctx context.Context, f uint64) (model.Project, error)
}
