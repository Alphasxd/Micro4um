package data

import (
	"github.com/Alphasxd/Micro4um/app/forum-check/internal/biz"
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	err := db.AutoMigrate(&biz.Check{})
	if err != nil {
		return err
	}
	return nil
}
