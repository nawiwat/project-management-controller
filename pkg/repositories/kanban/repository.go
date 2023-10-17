package kanban


import (
	"app-controller/pkg/model"
	"app-controller/pkg/repositories"
	"time"

	"gorm.io/gorm"

	"context"

	"github.com/pkg/errors"
)

type kanbanBoardRepository struct {
	db  *gorm.DB
	loc *time.Location
}

func New(db *gorm.DB) repositories.KanbanBoardRepository {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return &kanbanBoardRepository{
		db,
		loc,
	}
}

func (f *kanbanBoardRepository) Create(ctx context.Context, in model.Project) (error) {
	var kb model.KanbanBoard
	kb.ProjectID = in.ID
	if err := f.db.Create(&kb).Error; err != nil {
		return errors.Wrap(err, "fail to create kanban board")
	}

	return nil
}

func (f *kanbanBoardRepository) CreateColumn(ctx context.Context, in model.BoardColumn) (error) {
	if err := f.db.Create(&in).Error; err != nil {
		return errors.Wrap(err, "fail to create kanban board")
	}

	return nil
}

func (f *kanbanBoardRepository) Query(ctx context.Context, id uint64) ([]model.KanbanBoard, error) {
	var out []model.KanbanBoard
	err := f.db.Preload("Column").Where("project_id = ?", id).Find(&out).Error

	if err != nil {
		return nil, errors.Wrap(err, "failed to query kanban board")
	}

	return out, nil
}
