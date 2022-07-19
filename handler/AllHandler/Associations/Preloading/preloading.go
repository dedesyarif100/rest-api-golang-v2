package Preloading

import (
	"rest-api-golang-v2/handler/AllHandler/Associations/Preloading/Contents"
	"gorm.io/gorm"
)

func Preloading(db *gorm.DB, err any) {
	Contents.CustomPreloadingSQL(db, err)
	Contents.JoinsPreloading(db, err)
	Contents.NestedPreloading(db, err)
	Contents.Preload(db, err)
	Contents.PreloadAll(db, err)
	Contents.PreloadWithConditions(db, err)
}
