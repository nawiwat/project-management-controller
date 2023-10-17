package Users

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

func (f *projectsRepository) Create(ctx context.Context, in model.Project) (error) {
	if err := f.db.Create(&in).Error; err != nil {
		return errors.Wrap(err, "fail to create project")
	}

	return nil
}

func (f *projectsRepository) Query(ctx context.Context) ([]model.Project, error) {
	var out []model.Project
	err := f.db.Limit(50).Order("id asc").Find(&out).Error

	if err != nil {
		return nil, errors.Wrap(err, "failed to query project")
	}

	return out, nil
}

func (f *projectsRepository) QueryInfo(ctx context.Context, P_ID uint64) ([]model.Project, error) {
	var out []model.Project
	err := f.db.Preload("Membership.User").Where("id = ?",P_ID).Find(&out).Error

	if err != nil {
		return nil, errors.Wrap(err, "failed to query project")
	}

	return out, nil
}

func (f *projectsRepository) AddMember(ctx context.Context, in model.Membership) (error) {
	if err := f.db.Create(&in).Error; err != nil {
		return errors.Wrap(err, "fail to create project")
	}

	return nil
}