package contlr

import (
	"app-controller/pkg/model"
	"context"
	//"encoding/json"
)

func (s *service) GetProjects(ctx context.Context) ([]model.Project, error) {

	out, err := s.projectsRepo.Query(ctx)

	if err != nil {
		return []model.Project{}, err
	}

	return out, err
}

func (s *service) GetProjectInfo(ctx context.Context, f uint64) ([]model.Project, error) {

	out, err := s.projectsRepo.QueryInfo(ctx,f)

	if err != nil {
		return []model.Project{}, err
	}

	return out, err
}

func (s *service) AddProject(ctx context.Context, f model.Project) error {
	prj , err := 	s.projectsRepo.Create(ctx,f)

	if err != nil {
		return err
	}

	err = s.kanbanBoardRepo.Create(ctx,prj)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) AddMember(ctx context.Context, f model.Membership) error {
	err := 	s.projectsRepo.AddMember(ctx,f)

	if err != nil {
		return err
	}

	return nil
}
