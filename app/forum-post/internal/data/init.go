package data

import (
	"github.com/Alphasxd/Micro4um/app/forum-post/internal/biz"
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	err := db.AutoMigrate(&biz.Post{}, &biz.Plate{})
	if err != nil {
		return err
	}
	return nil
}
