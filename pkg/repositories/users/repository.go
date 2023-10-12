package Users

import (
	"app-controller/pkg/model"
	"app-controller/pkg/repositories"
	"time"

	"gorm.io/gorm"

	"context"

	"github.com/pkg/errors"
)

type usersRepository struct {
	db  *gorm.DB
	loc *time.Location
}

func New(db *gorm.DB) repositories.UsersRepository {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return &usersRepository{
		db,
		loc,
	}
}

func (f *usersRepository) Create(ctx context.Context, in model.User) (error) {
	if err := f.db.Create(&in).Error; err != nil {
		return errors.Wrap(err, "fail to create user")
	}

	return nil
}

func (f *usersRepository) Query(ctx context.Context) ([]model.User, error) {
	var out []model.User
	err := f.db.Limit(50).Order("id asc").Preload("Membership.Project").Find(&out).Error

	if err != nil {
		return nil, errors.Wrap(err, "failed to query users")
	}

	return out, nil
}
