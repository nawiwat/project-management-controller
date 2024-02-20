package projects

import (
	"app-controller/pkg/model"
	"app-controller/pkg/repositories"
	"time"

	"gorm.io/gorm"

	"context"

	"github.com/pkg/errors"
)

type projectsRepository struct {
	db  *gorm.DB
	loc *time.Location
}

func New(db *gorm.DB) repositories.ProjectsRepository {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return &projectsRepository{
		db,
		loc,
	}
}

func (f *projectsRepository) Create(ctx context.Context, in model.Project) (model.Project, error) {
	if err := f.db.Create(&in).Error; err != nil {
		return in, errors.Wrap(err, "fail to create project")
	}

	return in, nil
}

func (f *projectsRepository) Query(ctx context.Context, u string) ([]model.Project, error) {
	var out []model.Project
	var mem []model.Membership
	var pids []uint64

	err := f.db.Where("username = ?",u).Find(&mem).Error

	if err != nil {
		return nil, errors.Wrap(err, "failed to query membership")
	}

	for _,r := range(mem) {
		pids = append(pids, r.ProjectId)
	}

	tx := f.db.Where("")

	for _,r := range(pids) {
		tx.Or("id = ?",r)
	}
	
	tx.Find(&out)

	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "failed to query project")
	}

	return out, nil
}

func (f *projectsRepository) QueryInfo(ctx context.Context, pid uint64) (model.Project, error) {
	var out model.Project
	err := f.db.Preload("Task").Preload("Membership").Where("id = ?",pid).Find(&out).Error

	if err != nil {
		return model.Project{}, errors.Wrap(err, "failed to query project")
	}

	return out, nil
}

func (f *projectsRepository) AddMember(ctx context.Context, in model.Membership) (error) {
	if err := f.db.Create(&in).Error; err != nil {
		return errors.Wrap(err, "fail to create project")
	}

	return nil
}

func (f *projectsRepository) Update(ctx context.Context, in model.Project) (error) {
	if err := f.db.Updates(&in).Error; err != nil {
		return errors.Wrap(err, "fail to update project")
	}

	return nil
}
