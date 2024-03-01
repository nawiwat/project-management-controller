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
	err := f.db.Preload("Task").Preload("Membership").Preload("Invitation").Where("id = ?",pid).Find(&out).Error

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

func (f *projectsRepository) Delete(ctx context.Context, in uint64) (error) {
	var mem []model.Membership
	if err := f.db.Where("project_id = ?",in).Delete(&mem).Error; err != nil {
		return errors.Wrap(err, "fail to delete project")
	}

	return nil
}

func (f *projectsRepository) CreateInvite(ctx context.Context, in model.Invitation) (uint64 , error) {
	var mem []model.Membership
	var inv model.Invitation

	err := f.db.Where("project_id = ?",in.ProjectId).Find(&mem).Error

	if err != nil {
		return 0 , errors.Wrap(err, "failed to query membership")
	}

	for _,r := range(mem){
		if in.UserId == r.UserId {
			return 0 , errors.New("user already in the project")
		}
	}

	rf := f.db.Where("user_id = ?",in.UserId).Where("project_id = ?",in.ProjectId).Find(&inv).RowsAffected
	if rf != 0 {
		return 0 , errors.New("invite already exist")
	}

	if err := f.db.Create(&in).Error; err != nil {
		return 0 , errors.Wrap(err, "fail to create invitation")
	}

	if err := f.db.Where("user_id = ?",in.UserId).Where("project_id = ?",in.ProjectId).Find(&inv).Error; err != nil {
		return 0 , errors.Wrap(err, "fail to get invitation")
	}


	return inv.ID , nil
}

func (f *projectsRepository) DeleteInvite(ctx context.Context, in model.Invitation) (error) {
	if err := f.db.Delete(&in).Error; err != nil {
		return errors.Wrap(err, "fail to delete old invitation")
	}
	return nil
}

func (f *projectsRepository) GetInvite(ctx context.Context, in uint64) (model.Invitation , error) {
	var inv model.Invitation
	usrInvl := f.db.Where("id = ?",in).Find(&inv).RowsAffected
	if usrInvl == 0 {
		return model.Invitation{}, errors.New("invalid invitation")
	}

	return inv , nil
}