package migrations

import (
	"app-controller/pkg/model"

	"gorm.io/gorm"
)

// Migrate to migrate DB
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.User{},
		&model.Membership{},
		&model.Project{},
		&model.Notification{},
		&model.ProfileAttachment{},
		&model.BoardColumn{},
		&model.Invitation{},
		&model.UserToken{},
	); err != nil {
		return err
	}

	return nil
}
