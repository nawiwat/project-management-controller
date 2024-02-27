package tasks


import (
	"app-controller/pkg/model"
	"app-controller/pkg/repositories"
	"time"

	"gorm.io/gorm"

	"context"

	"github.com/pkg/errors"
)

type tasksRepository struct {
	db  *gorm.DB
	loc *time.Location
}

func New(db *gorm.DB) repositories.TasksRepository {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return &tasksRepository{
		db,
		loc,
	}
}

func (f *tasksRepository) Create(ctx context.Context, in model.Task ) ([]model.Task , error) {
	var out []model.Task
	
	if err := f.db.Create(&in).Error; err != nil {
		return nil , errors.Wrap(err, "fail to create task")
	}

	err := f.db.Preload("Attachments").Preload("Members").Preload("Comments").Preload("Kanban").Where("project_id = ?", in.ProjectId).Find(&out).Error

	if err != nil {
		return nil, errors.Wrap(err, "failed to query task")
	}

	return out , nil
}

func (f *tasksRepository) Query(ctx context.Context, id uint64 ) ([]model.Task , error) {
	var out []model.Task

	err := f.db.Preload("Attachments").Preload("Members").Preload("Comments").Preload("Kanban").Where("project_id = ?", id).Find(&out).Error

	if err != nil {
		return nil, errors.Wrap(err, "failed to query task")
	}

	return out , nil
}

func (f *tasksRepository) Update(ctx context.Context, in []model.Task ) ([]model.Task , error) {
	var out []model.Task
	for _,r := range(in) {
		if err := f.db.Updates(&r).Error; err != nil {
			return nil , errors.Wrap(err, "fail to update task")
		}
		for _,n := range(r.Attachments) {
			if err := f.db.Updates(n).Error; err != nil {
				return nil , errors.Wrap(err, "fail to update attachments")
			}
		}
		for _,n := range(r.Members) {
			if err := f.db.Updates(n).Error; err != nil {
				return nil , errors.Wrap(err, "fail to update members")
			}
		}
		for _,n := range(r.Comments) {
			if err := f.db.Updates(n).Error; err != nil {
				return nil , errors.Wrap(err, "fail to update comments")
			}
		}
	}
	

	err := f.db.Preload("Attachments").Preload("Members").Preload("Comments").Preload("Kanban").Where("project_id = ?", in[0].ProjectId).Find(&out).Error

	if err != nil {
		return nil, errors.Wrap(err, "failed to query task")
	}

	return out , nil
}