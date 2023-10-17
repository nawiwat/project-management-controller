package contlr

import (
	// "acw-crypto-risk-management/pkg/inputs/riskmgmt"
	"app-controller/pkg/model"
	"app-controller/pkg/repositories"
	"context"
	//"encoding/json"
)

type ControllerService interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUser(ctx context.Context, f uint64) ([]model.User, error)
	AddUser(ctx context.Context, f model.User) error

	GetProjects(ctx context.Context) ([]model.Project, error)
	GetProjectInfo(ctx context.Context, f uint64) ([]model.Project, error)
	AddProject(ctx context.Context, f model.Project) error
	AddMember(ctx context.Context, f model.Membership) error
}

type service struct {
	usersRepo         			repositories.UsersRepository
	projectsRepo 				repositories.ProjectsRepository
	// suspectedUserRepo 	   			repositories.SuspectedUserRepository
	// fraudRulesTransformer  			FraudRulesTransformer
	// freezeStatusRepo				repositories.FreezeStatusRepository
	// cache                  			CacheClient
}

func NewControllerService(
	usersRepo repositories.UsersRepository,
	projectsRepo repositories.ProjectsRepository,
	// suspectedUserRepo 	   repositories.SuspectedUserRepository,
	// freezeStatusRepo		repositories.FreezeStatusRepository,
	// cache CacheClient,
) ControllerService {
	baseService := &service{
		usersRepo:         usersRepo,
		projectsRepo: 		projectsRepo,
		// suspectedUserRepo:		suspectedUserRepo,
		// freezeStatusRepo:		freezeStatusRepo,
		// cache:                  cache,
	}

	return baseService
}