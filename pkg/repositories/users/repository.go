package users

import (
	"app-controller/pkg/model"
	"app-controller/pkg/repositories"
	"app-controller/pkg/utils"
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
	err := f.db.Preload("Membership").Preload("Notification").Preload("Attachment").Preload("TaskMember").Where("username = ?",username).Find(&out).Error

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

func (f *usersRepository) CreateNotification(ctx context.Context, in model.Notification) (error) {
	if err := f.db.Create(&in).Error; err != nil {
		return errors.Wrap(err, "fail to create notification")
	}
	return nil
}

func (f *usersRepository) GetNotification(ctx context.Context, in uint64) (model.Notification , error) {
	var out model.Notification
	usrInvl := f.db.Where("id = ?",in).Find(&out).RowsAffected
	if usrInvl == 0 {
		return model.Notification{} , errors.New("invalid notification")
	}
	return out , nil
}

func (f *usersRepository) DeleteNotification(ctx context.Context, in model.Notification) (error) {
	if err := f.db.Delete(&in).Error; err != nil {
		return errors.Wrap(err, "fail to delete old notification")
	}
	return  nil
}

func (f *usersRepository) UpdateNotification(ctx context.Context, in model.User , cur_task []model.Task) (error) {
	var usr model.User

	for _,r := range(in.Notification){
		switch r.Type {
			case "late", "critical", "close", "process":
				del := true
				for _,n := range(cur_task){
					if r.TaskId == n.ID {
						if n.Kanban.Column != "Done" {
							del = false
							break
						}
					}
				}
				if del {
					if err := f.db.Delete(&r).Error; err != nil {
						return errors.Wrap(err, "fail to delete old notification")
					}
				}
			}
	}

	for _,r := range(cur_task) {
		create := true
		if r.Kanban.Column == "Done" {
			create = false
		}
		for _,n := range(in.Notification){
			if r.ID == n.TaskId {
				create = false
			}
		}
		if create {
			noti := model.Notification{
			UserId: 		in.ID,
			SendBy:    		"",
			Type:  			"in_process",
			Description:    "",
			TaskId: 		r.ID,
			}
			if err := f.db.Create(&noti).Error; err != nil {
				return errors.Wrap(err, "fail to create new notification")
			}
		}
	}

	err := f.db.Where("id = ?",in.ID).Preload("Notification").Find(&usr).Error

	if err != nil {
		return errors.Wrap(err, "failed to query users stage:2")
	}

	for _,r := range(cur_task){
		status , err := utils.CheckDeadline(r)
		if err != nil {
			return errors.Wrap(err, "failed to check deadline")
		}

		if status == "in_process" {
			for _,n := range(usr.Notification) {
				if r.ID == n.TaskId {
					noti := model.Notification{
						ID: 			n.ID,
						UserId: 		n.UserId,
						SendBy:    		"",
						Type:  			"in_process",
						Description:    "",
						TaskId: 		n.TaskId,
						}
					if err := f.db.Updates(&noti).Error; err != nil {
						return errors.Wrap(err, "fail to update notification in process")
					}
					break
				}
			}
		} else if status == "close_due" {
			for _,n := range(usr.Notification) {
				if r.ID == n.TaskId {
					noti := model.Notification{
						ID: 			n.ID,
						UserId: 		n.UserId,
						SendBy:    		"Warning",
						Type:  			"close",
						Description:    r.Name + " deadline is 3 days left",
						TaskId: 		n.TaskId,
						}
					if err := f.db.Updates(&noti).Error; err != nil {
						return errors.Wrap(err, "fail to update notification close due")
					}
					break
				}
			}
		} else if status == "critical" {
			for _,n := range(usr.Notification) {
				if r.ID == n.TaskId {
					noti := model.Notification{
						ID: 			n.ID,
						UserId: 		n.UserId,
						SendBy:    		"Warning",
						Type:  			"critical",
						Description:    r.Name + " deadline is 1 day left",
						TaskId: 		n.TaskId,
						}
					if err := f.db.Updates(&noti).Error; err != nil {
						return errors.Wrap(err, "fail to update notification critical due")
					}
					break
				}
			}
		} else if status == "late" {
			for _,n := range(usr.Notification) {
				if r.ID == n.TaskId {
					noti := model.Notification{
						ID: 			n.ID,
						UserId: 		n.UserId,
						SendBy:    		"Warning",
						Type:  			"passed",
						Description:    r.Name + " has surpassed deadline",
						TaskId: 		n.TaskId,
						}
					if err := f.db.Updates(&noti).Error; err != nil {
						return errors.Wrap(err, "fail to update notification critical due")
					}
					break
				}
			}
		}
	}

	return  nil
}