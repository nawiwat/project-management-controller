package migrations

import (
	"app-controller/pkg/model"

	"gorm.io/gorm"
)

// Migrate to migrate DB
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.UserToken{},

		&model.User{},
		&model.ProfileAttachment{},
		&model.Notification{},
		
		&model.Project{},
		&model.Membership{},
		&model.Invitation{},
		
		&model.Task{},
		&model.Attachment{},
		&model.Comment{},
		&model.KanbanColumn{},
	); err != nil {
		return err
	}

	return nil
}
