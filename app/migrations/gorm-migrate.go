package migrations

import (
	_userData "taskApi/features/user/data"

	"gorm.io/gorm"
)

func InitDBMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{})
}
