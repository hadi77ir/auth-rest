package db

import (
	"auth-rest/internal/db/schema"
	"gorm.io/gorm"
)

func AutoMigrate(orm *gorm.DB) error {
	return orm.AutoMigrate(
		&schema.User{},
	)
}
