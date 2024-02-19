package contlr

import (
	"app-controller/pkg/model"
	"app-controller/pkg/repositories"
	"context"
	//"encoding/json"
)

type ControllerService interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUser(ctx context.Context, u string) (model.User, error)
	AddUser(ctx context.Context, f model.User) error
	EditUser(ctx context.Context, f model.User, u string) error
	EditProfile(ctx context.Context, f model.ProfileAttachment, u string) error
	Login(ctx context.Context, f model.UserLogin) (string, error)
	Auth(ctx context.Context, f model.UserLogin) (model.UserToken,error)

	GetProjects(ctx context.Context) ([]model.Project, error)
	GetProjectInfo(ctx context.Context, f uint64) (model.Project, error)
	AddProject(ctx context.Context, f model.Project, u string) error
	AddMember(ctx context.Context, f model.Membership) error

	AddBoardColumn(ctx context.Context, f model.BoardColumn) error
	
}

type service struct {
	usersRepo         			repositories.UsersRepository
	projectsRepo 				repositories.ProjectsRepository
	kanbanBoardRepo   			repositories.KanbanBoardRepository
}

func NewControllerService(
	usersRepo repositories.UsersRepository,
	projectsRepo repositories.ProjectsRepository,
	kanbanBoardRepo repositories.KanbanBoardRepository,
	
) ControllerService {
	baseService := &service{
		usersRepo:         usersRepo,
		projectsRepo: 		projectsRepo,
		kanbanBoardRepo:	kanbanBoardRepo,
	}

	return baseService
}