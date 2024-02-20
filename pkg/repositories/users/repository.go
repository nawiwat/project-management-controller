package users

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
	var usr model.User 
	usrdpl := f.db.Where("username = ?",in.Username).Find(&usr).RowsAffected
	if usrdpl != 0 {
		return errors.New("username duplicate")
	}
	if err := f.db.Create(&in).Error; err != nil {
		return errors.Wrap(err, "fail to create user")
	}

	err := f.db.Where("username = ?",in.Username).Find(&usr).RowsAffected
	if err == 0 {
		return errors.New("fail to get user from created")
	}

	in.Attachment.UserId = usr.ID

	if err := f.db.Create(&in.Attachment).Error; err != nil {
		return errors.Wrap(err, "fail to add attachment")
	}

	return nil
}

func (f *usersRepository) Query(ctx context.Context) ([]model.User, error) {
	var out []model.User
	err := f.db.Limit(50).Order("id asc").Find(&out).Error

	if err != nil {
		return nil, errors.Wrap(err, "failed to query users")
	}

	return out, nil
}

func (f *usersRepository) QueryInfo(ctx context.Context,username string) (model.User, error) {
	var out model.User
	err := f.db.Preload("Membership").Preload("Notification").Preload("Attachment").Where("username = ?",username).Find(&out).Error

	if err != nil {
		return model.User{}, errors.Wrap(err, "failed to query users")
	}

	return out, nil
}

func (f *usersRepository) QueryByUsername(ctx context.Context,username string) (model.User, error) {
	var out model.User
	usrInvl := f.db.Where("username = ?",username).Find(&out).RowsAffected

	if usrInvl == 0 {
		return model.User{} , errors.New("invalid Username")
	}

	err := f.db.Where("username = ?",username).Find(&out).Error

		
	if err != nil {
		return model.User{}, errors.Wrap(err, "failed to query users")
	}

	return out, nil
}

func (f *usersRepository) CreateToken(ctx context.Context,token model.UserToken) (error) {
	var userToken model.UserToken

	if utkdpl := f.db.Where("username = ?",token.Username).Find(&userToken).RowsAffected ; utkdpl != 0 {
		if err := f.db.Where("username = ?",token.Username).Delete(&token).Error; err != nil {
			return errors.Wrap(err, "fail to refresh token")
		}
	}

	if err := f.db.Create(&token).Error; err != nil {
		return errors.Wrap(err, "fail to create token")
	}

	return nil
}

func (f *usersRepository) QueryToken(ctx context.Context,username string) (model.UserToken, error) {
	var out model.UserToken
	usrInvl := f.db.Where("username = ?",username).Find(&out).RowsAffected

	if usrInvl == 0 {
		return model.UserToken{} , errors.New("invalid Username")
	}

	err := f.db.Where("username = ?",username).Find(&out).Error

	if err != nil {
		return model.UserToken{}, errors.Wrap(err, "failed to query token")
	}

	return out, nil
}

func (f *usersRepository) Update(ctx context.Context, in model.User) (error) {
	if err := f.db.Updates(&in).Error; err != nil {
		return errors.Wrap(err, "fail to update user")
	}
	return nil
}

func (f *usersRepository) UpdateProfile(ctx context.Context, in model.ProfileAttachment) (error) {
	if err := f.db.Where("id = ?",in.UserId).Updates(&in).Error; err != nil {
		return errors.Wrap(err, "fail to update profile")
	}
	return nil
}
