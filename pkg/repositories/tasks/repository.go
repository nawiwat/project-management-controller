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
	var cur_task model.Task
	for _,r := range(in) {
		if err := f.db.Save(&r).Error; err != nil {
			return nil , errors.Wrap(err, "fail to update task")
		}
		if err := f.db.Preload("Attachments").Preload("Members").Preload("Comments").Where("id = ?", r.ID).Find(&cur_task).Error; err != nil {
			return nil , errors.Wrap(err, "fail to find taskmember")
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

		for _, oldAttachment := range cur_task.Attachments {
			found := false
			for _, newAttachment := range r.Attachments {
				if oldAttachment.ID == newAttachment.ID {
					found = true
					break
				}
			}
			if !found {
				if err := f.db.Delete(&oldAttachment).Error; err != nil {
					return nil, errors.Wrap(err, "fail to delete old attachment")
				}
			}
		}

		for _, oldMember := range cur_task.Members {
			found := false
			for _, newMember := range r.Members {
				if oldMember.ID == newMember.ID {
					found = true
					break
				}
			}
			if !found {
				if err := f.db.Delete(&oldMember).Error; err != nil {
					return nil, errors.Wrap(err, "fail to delete old member")
				}
			}
		}

		for _, oldComment := range cur_task.Comments {
			found := false
			for _, newComment := range r.Comments {
				if oldComment.ID == newComment.ID {
					found = true
					break
				}
			}
			if !found {
				if err := f.db.Delete(&oldComment).Error; err != nil {
					return nil, errors.Wrap(err, "fail to delete old comment")
				}
			}
		}
	}
	

	err := f.db.Preload("Attachments").Preload("Members").Preload("Comments").Preload("Kanban").Where("project_id = ?", in[0].ProjectId).Find(&out).Error

	if err != nil {
		return nil, errors.Wrap(err, "failed to query task")
	}

	return out , nil
}

func (f *tasksRepository) Delete(ctx context.Context, in uint64 ) (error) {
	var cur_task model.Task

	if err := f.db.Preload("Attachments").Preload("Members").Preload("Comments").Preload("Kanban").Where("id = ?", in).Find(&cur_task).Error; err != nil {
		return errors.Wrap(err, "fail to find taskmember")
	}

	for _,r := range(cur_task.Attachments) {
		if err := f.db.Delete(&r).Error; err != nil {
			return errors.Wrap(err, "fail to delete attachment")
		}
	}

	for _,r := range(cur_task.Members) {
		if err := f.db.Delete(&r).Error; err != nil {
			return errors.Wrap(err, "fail to delete member")
		}
	}

	for _,r := range(cur_task.Comments) {
		if err := f.db.Delete(&r).Error; err != nil {
			return errors.Wrap(err, "fail to delete comment")
		}
	}

	if err := f.db.Where("task_id = ?",in).Delete(&cur_task.Kanban).Error; err != nil {
		return errors.Wrap(err, "fail to delete task column")
	}

	if err := f.db.Delete(&cur_task).Error; err != nil {
		return errors.Wrap(err, "fail to delete task")
	}

	return nil
}

func (f *tasksRepository) QueryByUserId(ctx context.Context, id uint64 ) ([]model.Task , error) {
	var mem []model.TaskMember
	var out []model.Task

	err := f.db.Where("user_id = ?", id).Find(&mem).Error

	if err != nil {
		return nil, errors.Wrap(err, "failed to query task member")
	}

	if len(mem) == 0 {
		return nil , nil
	}

	tx := f.db.Where("")

	for _,r := range(mem) {
		tx.Or("id = ?",r.TaskId)
	}
	
	tx.Preload("Kanban").Find(&out)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "failed to query task")
	}

	return out , nil
}