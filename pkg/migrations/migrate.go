package migrations

import (
	"app-controller/pkg/model"

	"gorm.io/gorm"
)

// Migrate to migrate DB
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		// &model.FraudRule{},
		&model.Membership{},
		&model.Project{},
		&model.User{},
	); err != nil {
		return err
	}

	return nil
}
