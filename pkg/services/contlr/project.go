package contlr

import (
	"app-controller/pkg/model"
	"context"
	//"encoding/json"
)

func (s *service) GetProjects(ctx context.Context, u string) ([]model.Project, error) {
	out, err := s.projectsRepo.Query(ctx, u)

	if err != nil {
		return []model.Project{}, err
	}

	return out, err
}

func (s *service) GetProjectInfo(ctx context.Context, f uint64) (model.Project, error) {

	out, err := s.projectsRepo.QueryInfo(ctx,f)

	if err != nil {
		return model.Project{}, err
	}

	return out, err
}

func (s *service) AddProject(ctx context.Context, f model.Project, u string) error {
	prj , err := 	s.projectsRepo.Create(ctx,f)

	if err != nil {
		return err
	}

	usr , err := s.usersRepo.QueryByUsername(ctx,u) 
	
	if err != nil {
		return err
	}

	pjOwner := model.Membership{
		ProjectId: prj.ID,
		UserId:    usr.ID,
		Username:  usr.Username,
		Role:      "Owner",
	}
	
	err = s.projectsRepo.AddMember(ctx,pjOwner)
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

func (s *service) EditProject(ctx context.Context, f model.Project) error {
	err := 	s.projectsRepo.Update(ctx, model.Project{
			ID: 			f.ID,	
			Name:   		f.Name,
			Email: 			f.Email,
			Budget: 		f.Budget,
			Deathline:		f.Deathline,	
			Github: 		f.Github,
			Phone: 			f.Phone,
			Description: 	f.Description,
	})

	if err != nil {
		return err
	}

	return nil
}