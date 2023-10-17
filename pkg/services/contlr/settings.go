package contlr

import (
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

	AddBoardColumn(ctx context.Context, f model.BoardColumn) error
	GetKanbanBoard(ctx context.Context, f uint64) ([]model.KanbanBoard, error)
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