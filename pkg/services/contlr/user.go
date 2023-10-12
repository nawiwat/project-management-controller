package contlr

import (
	// "acw-crypto-risk-management/pkg/inputs/riskmgmt"
	"app-controller/pkg/model"
	"context"
	//"encoding/json"
)

func (s *service) GetUsers(ctx context.Context) ([]model.User, error) {

	out, err := s.usersRepo.Query(ctx)

	if err != nil {
		return []model.User{}, err
	}

	return out, err
}

func (s *service) AddUser(ctx context.Context, f model.User) error {
	err := 	s.usersRepo.Create(ctx, model.User{
			Username:   	f.Username,
			Password:   	f.Password,
			Name:   		f.Name,
			Surname: 		f.Surname,
			Email: 			f.Email,
			Github: 		f.Github,
			Phone: 			f.Phone,
			Description: 	f.Description,
	})

	if err != nil {
		return err
	}

	return nil
}