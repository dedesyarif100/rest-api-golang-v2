package Scopes

import (
	"gorm.io/gorm"
	"rest-api-golang-v2/handler/AllHandler/Tutorials/Scopes/Contents"
)

func Performance(db *gorm.DB, err any) {
	Contents.DynamicallyTable(db, err)
	Contents.Query(db, err)
	Contents.Updates(db, err)
}