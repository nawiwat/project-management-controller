package contlr

import (
	"app-controller/pkg/model"
	"context"
	//"encoding/json"
)

func (s *service) CreateTask(ctx context.Context, t model.Task) ([]model.Task,error) {
	tsk , err := s.tasksRepo.Create(ctx,t)

	if err != nil {
		return nil , err
	}

	return tsk , nil
}

func (s *service) GetTask(ctx context.Context, id uint64) ([]model.Task,error) {
	tsk , err := s.tasksRepo.Query(ctx,id)

	if err != nil {
		return nil , err
	}

	return tsk , nil
}

func (s *service) UpdateTask(ctx context.Context, t []model.Task) ([]model.Task,error) {
	tsk , err := s.tasksRepo.Update(ctx,t)

	if err != nil {
		return nil , err
	}

	return tsk , nil
}

func (s *service) DeleteTask(ctx context.Context, id uint64) (error) {
	err := s.tasksRepo.Delete(ctx,id)

	if err != nil {
		return  err
	}

	return  nil
}