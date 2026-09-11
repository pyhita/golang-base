package dao

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	// 不停地加表
	return db.AutoMigrate(&User{})
}
