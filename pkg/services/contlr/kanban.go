package contlr

import (
	"app-controller/pkg/model"
	"context"
	//"encoding/json"
)

func (s *service) AddBoardColumn(ctx context.Context, f model.BoardColumn) error {
	err := s.kanbanBoardRepo.CreateColumn(ctx,f)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetKanbanBoard(ctx context.Context, f uint64) ([]model.KanbanBoard, error) {

	out, err := s.kanbanBoardRepo.Query(ctx,f)

	if err != nil {
		return []model.KanbanBoard{}, err
	}

	return out, err
}