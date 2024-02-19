package repositories

import (
	"app-controller/pkg/model"
	"context"
)

type  KanbanBoardRepository interface {
	//Create(ctx context.Context, in model.Project) (error)
	CreateColumn(ctx context.Context, in model.BoardColumn) (error)
	//Query(ctx context.Context, id uint64) ([]model.KanbanBoard, error)
}
